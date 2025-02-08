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
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request payload",
		})
		return
	}

	userExists, user, err := database.GetUserByUsername(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Error checking user existence",
		})
		return
	}
	if !userExists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User does not exist or the password is incorrect",
		})
		return
	}
	if user.UserPassword != creds.Password {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User does not exist or the password is incorrect",
		})
		return
	}

	token, err := helper.JwtHelper.IssueToken(map[string]interface{}{
		"username": user.Username,
		"role":     user.Role,
	}, "user", 60*60*24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("issue token err: %s", err.Error())})
		return
	}

	c.SetCookie("token", token, 60*60, "/api", "", false, true)

}
