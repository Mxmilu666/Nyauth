package handles

import (
	"net/http"
	"nyauth_backed/source/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

func NewResponse(status int, msg string, data interface{}) Response {
	return Response{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
}

func SendResponse(c *gin.Context, status int, msg string, data interface{}) {
	c.JSON(status, NewResponse(status, msg, data))
}

// JWTMiddleware 验证 JWT Gin 中间件
func JWTMiddleware(audience string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			SendResponse(c, http.StatusUnauthorized, "authorization header is missing", nil)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			SendResponse(c, http.StatusUnauthorized, "authorization header format must be Bearer {token}", nil)
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 验证 JWT
		token, err := helper.JwtHelper.VerifyToken(tokenString, audience)
		if err != nil {
			SendResponse(c, http.StatusUnauthorized, "invalid token", nil)
			c.Abort()
			return
		}

		// 将 claims 存储在上下文中
		c.Set("jwtClaims", token.Claims)
		c.Next()
	}
}
