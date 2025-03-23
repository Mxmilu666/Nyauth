package server

import (
	"net/http"
	"nyauth_backed/source/server/handles"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://ys.mihoyo.com")
	})

	api := r.Group("/api/v0")
	{
		api.GET("/code", func(c *gin.Context) {
			handles.SendResponse(c, http.StatusOK, "对不起，线路依然繁忙，请再等一下，或者稍后再打过来", nil)
		})

		// captcha
		api.GET("/captcha", handles.GetCaptcha)

		auth := api.Group("/account/auth")
		{
			auth.POST("/login", handles.UserLogin)
			auth.POST("/register", handles.UserRegister)
		}

		// 这个记得改成加个验证码
		api.POST("/account/sendcode", handles.SendVerificationCode)

		api.POST("/account/getaccountstatus", handles.GetAccountStatus)

		account := api.Group("/account", handles.JWTMiddleware("user"))
		{
			account.GET("/info", handles.UserInfo)
		}

		oauth := api.Group("/oauth", handles.JWTMiddleware("user"))
		{
			oauth.POST("/authorize", handles.OAuthAuthorize)
		}
	}
	return r
}
