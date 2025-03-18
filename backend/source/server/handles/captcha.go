package handles

import (
	"net/http"
	"nyauth_backed/source"

	"github.com/gin-gonic/gin"
)

// GetCaptcha 获取 captcha 验证码
func GetCaptcha(c *gin.Context) {
	turnstileSiteKey := source.AppConfig.Turnstile.SiteKey

	data := gin.H{
		"type": "turnstile",
		"id":   turnstileSiteKey,
	}

	SendResponse(c, http.StatusOK, "success", data)
}
