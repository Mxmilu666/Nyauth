package handles

import (
	"net/http"
	"nyauth_backed/source/database"
	"nyauth_backed/source/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// UserInfo 获取用户信息的处理函数
func UserInfo(c *gin.Context) {
	// 从上下文中获取用户ID
	claims, exists := c.Get("jwtClaims")
	if !exists {
		SendResponse(c, http.StatusUnauthorized, "JWT claims not found", nil)
		return
	}

	userID := claims.(jwt.MapClaims)["data"].(map[string]interface{})["user_id"].(string)

	// 通过用户ID获取用户信息
	user, err := database.GetUserByID(userID)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "Failed to get user info", nil)
		return
	}
	if user == nil {
		SendResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	// 输出格式
	userInfo := map[string]interface{}{
		"user_id":       user.UserID.Hex(),
		"user_uuid":     user.UserUUID,
		"user_name":     user.Username,
		"user_email":    user.UserEmail,
		"user_avatar":   user.Avatar,
		"role":          user.Role,
		"is_banned":     user.IsBanned,
		"register_at":   user.RegisterAt.Time().Format("2006-01-02 15:04:05"),
		"otp_enabled":   user.TOTPEnabled,
		"otp_enable_at": user.TOTPEnabledAt.Time().Format("2006-01-02 15:04:05"),
	}

	// 将用户信息封装在 user_info 对象中
	response := map[string]interface{}{
		"user_info": userInfo,
	}

	SendResponse(c, http.StatusOK, "success", response)
}

// UpdateUsername 更改用户名
func UpdateUsername(c *gin.Context) {
	var creds models.UpdateUsernameCredentials
	// 从上下文中获取用户ID
	claims, exists := c.Get("jwtClaims")
	if !exists {
		SendResponse(c, http.StatusUnauthorized, "JWT claims not found", nil)
		return
	}

	userID := claims.(jwt.MapClaims)["data"].(map[string]interface{})["user_id"].(string)

	if err := c.ShouldBindJSON(&creds); err != nil {
		SendResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// 检查新用户名是否为空
	if creds.Username == "" {
		SendResponse(c, http.StatusBadRequest, "Username cannot be empty", nil)
		return
	}

	// 检查新用户名是否已存在
	exists, _, err := database.GetUserByUsername(creds.Username)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "Failed to check username availability", nil)
		return
	}
	if exists {
		SendResponse(c, http.StatusConflict, "Username already exists", nil)
		return
	}

	// 更新用户信息
	updates := map[string]interface{}{
		"user_name": creds.Username,
	}

	err = database.UpdateUser(userID, updates)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "Failed to update username", nil)
		return
	}

	SendResponse(c, http.StatusOK, "Username updated successfully", nil)
}
