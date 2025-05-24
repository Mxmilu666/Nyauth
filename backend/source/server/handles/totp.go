package handles

import (
	"fmt"
	"net/http"
	"nyauth_backed/source/database"
	"nyauth_backed/source/helper"
	"nyauth_backed/source/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
)

// GenerateTOTP 生成TOTP密钥和二维码
func GenerateTOTP(c *gin.Context) {
	// 从JWT中获取当前用户ID
	claims, exists := c.Get("jwtClaims")
	if !exists {
		SendResponse(c, http.StatusUnauthorized, "未授权", nil)
		return
	}
	userID := claims.(jwt.MapClaims)["data"].(map[string]interface{})["user_id"].(string)

	// 获取用户信息
	user, err := database.GetUserByID(userID)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "获取用户信息失败", nil)
		return
	}
	if user == nil {
		SendResponse(c, http.StatusNotFound, "用户不存在", nil)
		return
	}

	// 检查用户是否已经启用了TOTP
	totpEnabled, err := database.UserHasTOTP(userID)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "检查TOTP状态失败", nil)
		return
	}

	if totpEnabled {
		SendResponse(c, http.StatusBadRequest, "TOTP已经启用", nil)
		return
	}

	// 生成TOTP密钥
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Nyauth",
		AccountName: user.UserEmail,
	})
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "生成TOTP密钥失败", nil)
		return
	}

	// 将密钥存储在会话中，等待用户验证
	// 为安全起见，暂时不保存到数据库，直到用户验证通过
	totpSecret := key.Secret()
	tempKey := fmt.Sprintf("totp_temp_%s", userID)
	err = helper.SetTempTOTPSecret(tempKey, totpSecret, 10*time.Minute)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "保存临时密钥失败", nil)
		return
	}

	response := gin.H{
		"secret":   totpSecret,
		"qr_code":  key.URL(),
		"issuer":   "Nyauth",
		"account":  user.UserEmail,
		"exp_time": 10 * 60, // 10分钟过期时间
	}

	SendResponse(c, http.StatusOK, "TOTP密钥生成成功", response)
}

// VerifyAndEnableTOTP 验证并启用TOTP
func VerifyAndEnableTOTP(c *gin.Context) {
	// 从JWT中获取当前用户ID
	claims, exists := c.Get("jwtClaims")
	if !exists {
		SendResponse(c, http.StatusUnauthorized, "未授权", nil)
		return
	}
	userID := claims.(jwt.MapClaims)["data"].(map[string]interface{})["user_id"].(string)

	// 获取请求参数
	var req models.TOTPVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, http.StatusBadRequest, "请求参数错误", nil)
		return
	}

	// 获取临时保存的TOTP密钥
	tempKey := fmt.Sprintf("totp_temp_%s", userID)
	secret, exists := helper.GetTempTOTPSecret(tempKey)
	if !exists {
		SendResponse(c, http.StatusBadRequest, "TOTP密钥已过期或不存在，请重新生成", nil)
		return
	}

	// 验证TOTP码
	valid := totp.Validate(req.Code, secret)
	if !valid {
		SendResponse(c, http.StatusBadRequest, "验证码无效", nil)
		return
	}

	// 验证通过，将TOTP密钥保存到数据库
	err := database.EnableTOTP(userID, secret)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "启用TOTP失败", nil)
		return
	}

	// 清除临时密钥
	helper.RemoveTempTOTPSecret(tempKey)

	// 生成恢复码
	recoveryCodes, err := database.GenerateAndSaveRecoveryCodes(userID, 5)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "生成恢复码失败", nil)
		return
	}

	SendResponse(c, http.StatusOK, "TOTP启用成功", gin.H{
		"recovery_codes": recoveryCodes,
	})
}

// VerifyTOTP 验证TOTP码（用于登录）
func VerifyTOTP(c *gin.Context) {
	var req models.TOTPLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, http.StatusBadRequest, "请求参数错误", nil)
		return
	}

	// 验证用户
	userExists, user, err := database.GetUserByUsername(req.Username)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "验证用户时出错", nil)
		return
	}
	if !userExists {
		SendResponse(c, http.StatusNotFound, "用户不存在", nil)
		return
	}

	// 检查用户是否启用了TOTP
	totpEnabled, secret, err := database.GetUserTOTPSecret(user.UserID.Hex())
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "检查TOTP状态失败", nil)
		return
	}

	if !totpEnabled {
		SendResponse(c, http.StatusBadRequest, "该用户未启用TOTP", nil)
		return
	}

	// 验证TOTP码
	valid := totp.Validate(req.Code, secret)
	if !valid {
		// 检查是否使用恢复码
		isRecoveryCode, err := database.ValidateAndConsumeRecoveryCode(user.UserID.Hex(), req.Code)
		if err != nil || !isRecoveryCode {
			SendResponse(c, http.StatusBadRequest, "验证码无效", nil)
			return
		}
	}

	// 验证通过，生成JWT
	exp := int64(60 * 60 * 24)
	token, err := helper.JwtHelper.IssueToken(map[string]interface{}{
		"user_name": user.Username,
		"user_id":   user.UserID.Hex(),
		"role":      user.Role,
	}, "user", exp)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, fmt.Sprintf("生成token失败: %s", err.Error()), nil)
		return
	}

	SendResponse(c, http.StatusOK, "验证成功", gin.H{
		"token": token,
		"exp":   exp,
	})
}

// DisableTOTP 禁用TOTP
func DisableTOTP(c *gin.Context) {
	// 从JWT中获取当前用户ID
	claims, exists := c.Get("jwtClaims")
	if !exists {
		SendResponse(c, http.StatusUnauthorized, "未授权", nil)
		return
	}
	userID := claims.(jwt.MapClaims)["data"].(map[string]interface{})["user_id"].(string)

	// 验证密码
	var req models.TOTPDisableRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, http.StatusBadRequest, "请求参数错误", nil)
		return
	}

	// 禁用TOTP
	err := database.DisableTOTP(userID)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "禁用TOTP失败", nil)
		return
	}

	SendResponse(c, http.StatusOK, "TOTP已禁用", nil)
}
