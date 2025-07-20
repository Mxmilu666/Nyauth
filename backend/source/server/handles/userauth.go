package handles

import (
	"errors"
	"fmt"
	"net/http"
	"nyauth_backed/source/database"
	"nyauth_backed/source/helper"
	"nyauth_backed/source/logger"
	"nyauth_backed/source/models"
	"nyauth_backed/source/untils"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

// 用户登录
func UserLogin(c *gin.Context) {
	var creds models.LoginCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		SendResponse(c, http.StatusBadRequest, "请求无效", nil)
		return
	}

	// 验证Turnstile验证码
	if creds.Secretkey == "" {
		SendResponse(c, http.StatusBadRequest, "验证码不能为空", nil)
		return
	}

	success, err := VerifyTurnstile(creds.Secretkey)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "验证码验证过程出错", nil)
		return
	}
	if !success {
		SendResponse(c, http.StatusBadRequest, "验证码验证失败", nil)
		return
	}

	userExists, user, err := database.GetUserByUsername(creds.Username)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "不..不好..数据库坏掉了❤", nil)
		return
	}
	if !userExists || user.UserPassword != creds.Password {
		SendResponse(c, http.StatusNotFound, "用户不存在或密码不正确", nil)
		return
	}

	// 检查用户是否启用了TOTP
	totpEnabled, err := database.UserHasTOTP(user.UserID.Hex())
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "检查TOTP状态失败", nil)
		return
	}

	// 如果用户启用了TOTP
	if totpEnabled {
		// 如果提供了TOTP代码，直接验证
		if creds.TotpCode != "" {
			// 获取用户TOTP密钥
			_, secret, err := database.GetUserTOTPSecret(user.UserID.Hex())
			if err != nil {
				SendResponse(c, http.StatusInternalServerError, "获取TOTP密钥失败", nil)
				return
			}

			// 验证TOTP代码
			valid := totp.Validate(creds.TotpCode, secret)
			if !valid {
				// 检查是否使用恢复码
				isRecoveryCode, err := database.ValidateAndConsumeRecoveryCode(user.UserID.Hex(), creds.TotpCode)
				if err != nil || !isRecoveryCode {
					SendResponse(c, http.StatusBadRequest, "TOTP验证码无效", nil)
					return
				}
			}

			// TOTP验证通过，继续生成token
		} else {
			// 未提供TOTP代码，返回需要TOTP验证的响应
			SendResponse(c, http.StatusOK, "需要TOTP验证", gin.H{
				"require_totp": true,
				"username":     user.Username,
			})
			return
		}
	}

	exp := int64(60 * 60 * 24)

	token, err := helper.JwtHelper.IssueToken(map[string]interface{}{
		"user_name": user.Username,
		"user_id":   user.UserID,
		"role":      user.Role,
	}, "user", exp)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, fmt.Sprintf("issue token err: %s", err.Error()), nil)
		return
	}

	SendResponse(c, http.StatusOK, "获取 Token 成功", gin.H{
		"token": token,
		"exp":   exp,
	})
}

// 用户注册
func UserRegister(c *gin.Context) {
	var creds models.RegisterCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		SendResponse(c, http.StatusBadRequest, "请求无效", nil)
		return
	}

	// 验证临时注册码
	if !helper.VerifyTempCode(creds.Useremail, creds.TempCode, "register") {
		SendResponse(c, http.StatusBadRequest, "验证已过期或无效，请重新验证邮箱", nil)
		return
	}

	userExists, _, err := database.GetUserByUsername(creds.Username)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "不..不好..数据库坏掉了❤", nil)
		return
	}
	if userExists {
		SendResponse(c, http.StatusConflict, "邮箱或用户名已存在", nil)
		return
	}

	avatar := "https://cravatar.cn/avatar/" + untils.MD5(creds.Useremail) + "?s=256"

	userId, err := database.CreateUser(creds.Username, creds.Useremail, creds.Password, avatar)
	if err != nil {
		logger.Error("Failed to create user: ", err)
		SendResponse(c, http.StatusInternalServerError, "创建用户时出错", nil)
		return
	}

	// 生成 JWT token
	exp := int64(60 * 60 * 24)
	token, err := helper.JwtHelper.IssueToken(map[string]interface{}{
		"user_name": creds.Username,
		"user_id":   userId,
		"role":      "user",
	}, "user", exp)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, fmt.Sprintf("issue token err: %s", err.Error()), nil)
		return
	}

	SendResponse(c, http.StatusOK, "注册成功", gin.H{
		"token": token,
		"exp":   exp,
	})
}

func SendVerificationCode(c *gin.Context) {
	var creds models.EmailCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		SendResponse(c, http.StatusBadRequest, "请求无效", nil)
		return
	}

	// 验证Turnstile验证码
	//if creds.Secretkey == "" {
	//	SendResponse(c, http.StatusBadRequest, "验证码不能为空", nil)
	//	return
	//}

	//success, err := VerifyTurnstile(creds.Secretkey)
	//if err != nil {
	//	SendResponse(c, http.StatusInternalServerError, "验证码验证过程出错", nil)
	//	return
	//}
	//if !success {
	//	SendResponse(c, http.StatusBadRequest, "验证码验证失败", nil)
	//	return
	//}

	// 从查询参数中读取 usefor
	usefor := c.Query("usefor")
	if usefor == "" {
		SendResponse(c, http.StatusBadRequest, "请求无效", nil)
		return
	}

	// 根据 usefor 执行对应的操作
	switch usefor {
	case "register", "reset_password", "multi_identity":
		err := helper.SendVerificationCodeByEmail(creds.Useremail, usefor)
		if err != nil {
			if errors.Is(err, helper.ErrVerificationCodeExists) {
				// 验证码已存在并且未过期
				SendResponse(c, http.StatusTooManyRequests, "验证码还未过期呢，请等待一会再发送吧", nil)
				return
			}
			logger.Error("Failed to send verification code: ", err)
			SendResponse(c, http.StatusInternalServerError, "发送验证码时出错", nil)
			return
		}
		SendResponse(c, http.StatusOK, "发送验证码成功! 请注意查收~", nil)
	default:
		SendResponse(c, http.StatusBadRequest, "无效的参数", nil)
	}
}

// GetAccountStatus 检查用户是否存在
func GetAccountStatus(c *gin.Context) {
	var query models.GetAccountStatusCredentials

	if err := c.ShouldBindJSON(&query); err != nil {
		SendResponse(c, http.StatusBadRequest, "请求无效", nil)
		return
	}

	// 查找是否存在
	userExists, user, err := database.GetUserByUsername(query.Username)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "检查用户存在时出错", nil)
		return
	}

	if userExists {
		// 用户存在
		SendResponse(c, http.StatusOK, "success", gin.H{
			"exists": true,
			"user_info": map[string]interface{}{
				"email":       user.UserEmail,
				"enable_totp": user.TOTPEnabled,
			},
		})
	} else {
		// 用户不存在
		SendResponse(c, http.StatusOK, "success", gin.H{
			"exists": false,
		})
	}
}

// 验证验证码并生成临时注册码
func VerifyEmailCode(c *gin.Context) {
	var creds models.VerifyCodeCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		SendResponse(c, http.StatusBadRequest, "笨蛋！你..你放了什么东西进来❤", nil)
		return
	}

	// 从查询参数中读取 usefor
	usefor := c.Query("usefor")
	if usefor == "" {
		SendResponse(c, http.StatusBadRequest, "笨蛋！你..你放了什么东西进来❤", nil)
		return
	}

	// 验证验证码
	if !helper.VerifyCode(creds.Useremail, creds.Code, usefor) {
		SendResponse(c, http.StatusBadRequest, "验证码错误或已过期", nil)
		return
	}

	// 生成临时注册码，30分钟有效期
	tempCode, err := helper.GenerateTempCode(creds.Useremail, usefor, 30)
	if err != nil {
		logger.Error("Failed to generate temporary code: ", err)
		SendResponse(c, http.StatusInternalServerError, "不..不行❤里面坏掉了..❤", nil)
		return
	}

	SendResponse(c, http.StatusOK, "success", gin.H{
		"temp_code": tempCode,
		"exp":       30 * 60,
	})
}
