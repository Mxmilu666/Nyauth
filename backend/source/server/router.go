package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"nyauth_backed/source/server/handles"
	"strings"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) *gin.Engine {

	target, _ := url.Parse("http://localhost:5173")
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 设置反向代理中间件 - 对除了 /api/v0 外的所有路由
	r.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/v0") || strings.HasPrefix(c.Request.URL.Path, "/.well-known") {
			c.Next()
			return
		}

		// 其他所有路由反向代理到 localhost:5173
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort() // 阻止后续处理
	})

	///r.GET("/", func(c *gin.Context) {
	///	c.Redirect(http.StatusFound, "https://ys.mihoyo.com")
	///})

	r.GET("/.well-known/openid-configuration", handles.GetOpenIDConfiguration)
	r.GET("/.well-known/jwks.json", handles.GetJWKS)

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
				multiAccount.GET("/info", handles.GetMultiIdentities)
				// 创建多用户身份
				multiAccount.POST("/create", handles.CreateMultiIdentity)
			}
		}

		oauth := api.Group("/oauth")
		{
			oauthProtected := oauth.Group("", handles.JWTMiddleware("user"))
			{
				oauthProtected.GET("/authorize", handles.OAuthAuthorize)
				oauthProtected.POST("/getclientinfo", handles.GetClientinfo)
			}

			oauth.POST("/token", handles.OAuthToken)
		}
	}
	return r
}
