package handles

import (
	"fmt"
	"net/http"
	"nyauth_backed/source/database"
	"nyauth_backed/source/helper"
	"nyauth_backed/source/models"

	"github.com/gin-gonic/gin"
)

// 用户登录
func UserLogin(c *gin.Context) {
	var creds models.LoginCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		SendResponse(c, http.StatusBadRequest, "请求负载无效", nil)
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
		"username": user.Username,
		"role":     user.Role,
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

}

func SendVerificationCode(c *gin.Context) {
	var creds models.EmailCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		SendResponse(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	err := helper.SendVerificationCodeByEmail(creds.Useremail)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "Failed to send verification code", nil)
		return
	}

	SendResponse(c, http.StatusOK, "Verification code sent successfully", nil)
}
