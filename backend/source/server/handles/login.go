package handles

import (
	"fmt"
	"net/http"
	"nyauth_backed/source/database"
	"nyauth_backed/source/helper"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Secretkey string `json:"turnstile_secretkey"`
}

func Userlogin(c *gin.Context) {
	var creds Credentials
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
