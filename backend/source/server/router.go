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

		api.POST("/account/sendcode", handles.SendVerificationCode)

		api.POST("/account/verifycode", handles.VerifyEmailCode)

		api.POST("/account/getaccountstatus", handles.GetAccountStatus)

		account := api.Group("/account", handles.JWTMiddleware("user"))
		{
			// 获取用户信息
			account.GET("/info", handles.UserInfo)
			// 修改用户名
			account.POST("/update/username", handles.UpdateUsername)

			// 多用户
			multiAccount := account.Group("/multi")
			{
				// 获取多用户信息
				multiAccount.GET("/info", handles.OAuthAuthorize)
			}
		}

		oauth := api.Group("/oauth", handles.JWTMiddleware("user"))
		{
			oauth.POST("/authorize", handles.OAuthAuthorize)
		}
	}
	return r
}
