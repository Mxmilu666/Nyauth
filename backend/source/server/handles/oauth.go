package handles

import (
	"fmt"
	"net/http"
	"nyauth_backed/source/database"
	"nyauth_backed/source/helper"
	"nyauth_backed/source/models"
	"nyauth_backed/source/oauth"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func OAuthAuthorize(c *gin.Context) {
	// 从请求中获取OAuth参数
	clientID := c.Query("client_id")
	redirectURI := c.Query("redirect_uri")
	responseType := c.Query("response_type")
	scope := c.Query("scope")
	state := c.Query("state") // 可选参数，用于防止CSRF攻击

	// 验证必要参数
	if clientID == "" || redirectURI == "" || responseType == "" {
		SendResponse(c, http.StatusBadRequest, "参数不完整", nil)
		return
	}

	// 验证response_type，目前只支持code模式
	if responseType != "code" {
		SendResponse(c, http.StatusBadRequest, "不支持的 response_type", nil)
		return
	}

	// 获取客户端信息
	client, err := database.GetClientByClientID(clientID)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "服务器错误，无法验证客户端信息", nil)
		fmt.Printf("GetClientByClientID err: %s\n", err.Error())
		return
	}

	if client == nil {
		SendResponse(c, http.StatusNotFound, "未找到指定的应用", nil)
		return
	}

	// 验证重定向URI是否与注册的一致
	if client.RedirectURI != redirectURI {
		SendResponse(c, http.StatusBadRequest, "提供的 redirect_uri 不匹配", nil)
		return
	}

	// 检查scope中是否包含openid
	if scope == "" || !strings.Contains(scope, "openid") {
		SendResponse(c, http.StatusBadRequest, "缺少必要的 openid 权限范围", nil)
		return
	}

	// 从上下文中获取用户ID
	claims, exists := c.Get("jwtClaims")
	if !exists {
		SendResponse(c, http.StatusUnauthorized, "JWT claims not found", nil)
		return
	}

	userID := claims.(jwt.MapClaims)["data"].(map[string]interface{})["user_id"].(string)

	// 生成授权码
	authCode, err := oauth.CreateAuthorizationCode(clientID, userID)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "生成授权码失败", nil)
		fmt.Printf("CreateAuthorizationCode err: %s\n", err.Error())
		return
	}

	// 构建重定向URL
	var redirectURL string
	if strings.Contains(redirectURI, "?") {
		// 如果重定向URI已经包含查询参数，使用&连接
		redirectURL = fmt.Sprintf("%s&code=%s", redirectURI, authCode)
	} else {
		// 否则使用?开始查询参数
		redirectURL = fmt.Sprintf("%s?code=%s", redirectURI, authCode)
	}

	if state != "" {
		redirectURL = fmt.Sprintf("%s&state=%s", redirectURL, state)
	}

	// 返回URL给前端，由前端进行跳转，而不是直接重定向
	SendResponse(c, http.StatusOK, "授权成功", gin.H{
		"redirect_url": redirectURL,
	})
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
		SendResponse(c, http.StatusNotFound, "没有这个 client 哦", nil)
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
		"description": client.Description,
		"created_by":  creator.Username,
		"avatar":      client.Avatar,
		"status":      client.Status,
		"created_at":  client.CreatedAt,
		"permissions": client.Permissions,
	}

	SendResponse(c, http.StatusOK, "获取应用信息成功", clientInfo)
}

func OAuthToken(c *gin.Context) {
	// 解析表单数据
	grantType := c.PostForm("grant_type")
	code := c.PostForm("code")
	clientID := c.PostForm("client_id")
	clientSecret := c.PostForm("client_secret")
	redirectURI := c.PostForm("redirect_uri")
	nonce := c.PostForm("nonce")

	// 验证必要参数
	if grantType == "" || code == "" || clientID == "" || clientSecret == "" || redirectURI == "" {
		SendResponse(c, http.StatusBadRequest, "参数不完整", nil)
		return
	}

	// 验证授权类型，目前只支持authorization_code
	if grantType != "authorization_code" {
		SendResponse(c, http.StatusBadRequest, "不支持的 grant_type", nil)
		return
	}

	// 验证客户端信息
	client, err := database.GetClientByClientID(clientID)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "无法验证客户端信息，等一下再试吧~", nil)
		fmt.Printf("GetClientByClientID err: %s\n", err.Error())
		return
	}

	if client == nil {
		SendResponse(c, http.StatusNotFound, "没有这个 client 哦", nil)
		return
	}

	// 验证客户端密钥
	if client.ClientSecret != clientSecret {
		SendResponse(c, http.StatusUnauthorized, "密钥不正确", nil)
		return
	}

	// 验证重定向URI是否与注册的一致
	fmt.Printf("RedirectURI: %s\n", redirectURI)
	if client.RedirectURI != redirectURI {
		SendResponse(c, http.StatusBadRequest, "提供的 redirect_uri 不匹配", nil)
		return
	}

	// 验证授权码并获取相关信息
	authInfo, valid := oauth.GetAuthorizationCode(code)
	if !valid {
		SendResponse(c, http.StatusBadRequest, "授权码无效或已过期", nil)
		return
	}

	// 生成访问令牌
	accessToken, err := oauth.CreateToken(clientID, authInfo.UserID, client.Permissions)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "生成访问令牌失败", nil)
		fmt.Printf("CreateToken err: %s\n", err.Error())
		return
	}

	// 获取当前的token对象，以便返回过期时间
	tokenObj, exists := oauth.GetToken(accessToken)
	if !exists {
		SendResponse(c, http.StatusInternalServerError, "获取令牌信息失败", nil)
		return
	}

	idToken, err := helper.JwtHelper.IssueOIDCToken(authInfo.UserID, clientID, nonce, 6400)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "生成ID令牌失败", nil)
		fmt.Printf("IssueOIDCToken err: %s\n", err.Error())
		return
	}

	// 返回访问令牌
	expiresIn := int(time.Until(tokenObj.Exp).Seconds())
	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"expires_in":   expiresIn,
		"scope":        tokenObj.Scope,
		"id_token":     idToken,
	})
}
