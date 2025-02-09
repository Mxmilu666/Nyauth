package server

import (
	"fmt"
	"log"
	"net/url"
	"nyauth_backed/source"
	"nyauth_backed/source/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Setupserver() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(filterLogs())
	r.Use(corsMiddleware())
	r = initRouter(r)

	// start http server
	address := fmt.Sprintf("%s:%d", source.AppConfig.Server.Host, source.AppConfig.Server.Port)
	logger.Info(fmt.Sprintf("Starting server on %s\n", address))
	if err := r.Run(address); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func filterLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		fullURL, err := url.QueryUnescape(c.Request.URL.String())
		if err != nil {
			logger.Error("Error decoding URL:", err)
			return
		}

		fullURL = strings.ReplaceAll(fullURL, "%", "%%")

		c.Next()

		latency := time.Since(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		userAgent := c.Request.UserAgent()

		// 构造日志输出格式
		logger.Info(fmt.Sprintf(
			"%3d | %13v | %15s | %-7s | %s | %s\n",
			statusCode, // 状态码
			latency,    // 延迟时间
			clientIP,   // 客户端IP
			method,     // 请求方法
			userAgent,  // 用户代理
			fullURL,    // 完整的 URL 包括路径和查询参数
		))
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
