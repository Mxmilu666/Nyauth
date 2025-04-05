package handles

import (
	"fmt"
	"net/http"
	"nyauth_backed/source/database"
	"nyauth_backed/source/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func OAuthAuthorize(c *gin.Context) {
	SendResponse(c, http.StatusBadRequest, "你说得对但是后面忘了", nil)
}

func GetClientinfo(c *gin.Context) {
	// 从请求体中获取客户端ID
	var creds models.GetClientinfoCredentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		SendResponse(c, http.StatusBadRequest, "笨蛋！你..你放了什么东西进来❤", nil)
		return
	}

	// 检查客户端ID格式是否有效
	_, err := bson.ObjectIDFromHex(creds.ClientID)
	if err != nil {
		SendResponse(c, http.StatusBadRequest, "笨蛋！你..你放了什么东西进来❤", nil)
		return
	}

	// 获取客户端信息
	client, err := database.GetClientByClientID(creds.ClientID)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "不..不行❤里面坏掉了..❤", nil)
		fmt.Printf("GetClientByClientID err: %s\n", err.Error())
		return
	}

	if client == nil {
		SendResponse(c, http.StatusNotFound, "没有这个client哦", nil)
		return
	}

	// 获取创建者的用户信息
	creator, err := database.GetUserByID(client.CreatedBy)
	if err != nil || creator == nil {
		SendResponse(c, http.StatusNotFound, "啊嘞，怎么没有找到创建者?", nil)
		return
	}

	// 构造返回数据，移除敏感信息
	clientInfo := gin.H{
		"client_id":   client.ID.Hex(),
		"client_name": client.ClientName,
		"created_by":  creator.Username,
		"avatar":      client.Avatar,
		"permissions": client.Permissions,
		"status":      client.Status,
		"created_at":  client.CreatedAt,
	}

	SendResponse(c, http.StatusOK, "获取应用信息成功", clientInfo)
}
