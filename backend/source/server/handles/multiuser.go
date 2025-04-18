package handles

import (
	"fmt"
	"net/http"
	"nyauth_backed/source/database"
	"nyauth_backed/source/helper"
	"nyauth_backed/source/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// CreateMultiIdentity 创建用户多身份
func CreateMultiIdentity(c *gin.Context) {
	// 从JWT中获取当前用户ID
	claims, exists := c.Get("jwtClaims")
	if !exists {
		SendResponse(c, http.StatusUnauthorized, "未授权", nil)
		return
	}
	userID := claims.(jwt.MapClaims)["data"].(map[string]interface{})["user_id"].(string)

	var req models.MultiuserCredentials

	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, http.StatusBadRequest, "请求参数错误", nil)
		return
	}

	// 验证临时码
	if !helper.VerifyTempCode(req.Email, req.TempCode, "multi_identity") {
		SendResponse(c, http.StatusBadRequest, "验证已过期或无效，请重新验证邮箱", nil)
		return
	}

	// 创建新的身份
	identityID, err := database.CreateUserIdentity(userID, req.Email, req.DisplayName, req.Description, req.Avatar)
	if err != nil {
		fmt.Printf("CreateUserIdentity err: %s\n", err.Error())
		SendResponse(c, http.StatusInternalServerError, "创建多身份失败", nil)
		return
	}

	SendResponse(c, http.StatusOK, "创建多身份成功", gin.H{
		"identity_id": identityID,
	})
}

// GetMultiIdentities 获取用户所有多身份
func GetMultiIdentities(c *gin.Context) {
	// 从JWT中获取当前用户ID
	claims, exists := c.Get("jwtClaims")
	if !exists {
		SendResponse(c, http.StatusUnauthorized, "未授权", nil)
		return
	}
	userID := claims.(jwt.MapClaims)["data"].(map[string]interface{})["user_id"].(string)

	// 获取用户的所有多身份
	identities, err := database.GetUserIdentities(userID)
	if err != nil {
		fmt.Printf("GetUserIdentities err: %s\n", err.Error())
		SendResponse(c, http.StatusInternalServerError, "获取多身份列表失败", nil)
		return
	}

	SendResponse(c, http.StatusOK, "获取多身份列表成功", gin.H{
		"identities": identities,
	})
}
