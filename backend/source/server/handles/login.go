package handles

import (
	"net/http"
	"nyauth_backed/source/database"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userExists, err := database.CheckUserExists(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user existence"})
		return
	}
	if !userExists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User exists"})
}
