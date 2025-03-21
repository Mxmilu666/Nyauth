package handles

import (
	"fmt"
	"net/http"
	"nyauth_backed/source/database"
	"nyauth_backed/source/helper"
	"nyauth_backed/source/logger"
	"nyauth_backed/source/models"

	"github.com/gin-gonic/gin"
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
		SendResponse(c, http.StatusInternalServerError, "检查用户存在时出错，数据库爆炸啦！", nil)
		return
	}
	if !userExists || user.UserPassword != creds.Password {
		SendResponse(c, http.StatusNotFound, "用户不存在或密码不正确", nil)
		return
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

	if !helper.VerifyCode(creds.Useremail, creds.Code, "register") {
		SendResponse(c, http.StatusBadRequest, "验证码错误", nil)
		return
	}

	userExists, _, err := database.GetUserByUsername(creds.Username)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "检查用户存在时出错，数据库爆炸啦！", nil)
		return
	}
	if userExists {
		SendResponse(c, http.StatusConflict, "邮箱或用户名已存在", nil)
		return
	}

	avatar := "https://cravatar.cn/avatar/" + MD5(creds.Useremail) + "?d=identicon"

	err = database.CreateUser(creds.Username, creds.Useremail, creds.Password, avatar)
	if err != nil {
		logger.Error("Failed to create user: ", err)
		SendResponse(c, http.StatusInternalServerError, "创建用户时出错", nil)
		return
	}

	SendResponse(c, http.StatusOK, "用户注册成功", nil)
}

func SendVerificationCode(c *gin.Context) {
	var creds models.EmailCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		SendResponse(c, http.StatusBadRequest, "请求无效", nil)
		return
	}

	// 从查询参数中读取 usefor
	usefor := c.Query("usefor")
	if usefor == "" {
		SendResponse(c, http.StatusBadRequest, "请求无效", nil)
		return
	}

	// 根据 usefor 执行对应的操作
	switch usefor {
	case "register":
		err := helper.SendVerificationCodeByEmail(creds.Useremail, "register")
		if err != nil {
			logger.Error("Failed to send verification code: ", err)
			SendResponse(c, http.StatusInternalServerError, "发送验证码时出错", nil)
			return
		}
		SendResponse(c, http.StatusOK, "发送验证码成功! 请注意查收~", nil)
	case "reset_password":
		err := helper.SendVerificationCodeByEmail(creds.Useremail, "reset_password")
		if err != nil {
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
			"user_info": map[string]string{
				"email": user.UserEmail,
			},
		})
	} else {
		// 用户不存在
		SendResponse(c, http.StatusOK, "success", gin.H{
			"exists": false,
		})
	}
}
