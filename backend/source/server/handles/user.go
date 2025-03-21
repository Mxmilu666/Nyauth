package handles

import (
	"net/http"
	"nyauth_backed/source/database"

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
		"user_id":     user.UserID.Hex(),
		"user_name":   user.Username,
		"user_email":  user.UserEmail,
		"user_avatar": user.Avatar,
		"role":        user.Role,
		"is_banned":   user.IsBanned,
		"register_at": user.RegisterAt,
	}

	// 将用户信息封装在 user_info 对象中
	response := map[string]interface{}{
		"user_info": userInfo,
	}

	SendResponse(c, http.StatusOK, "success", response)
}
