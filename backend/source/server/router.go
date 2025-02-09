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

		account := api.Group("/account")
		{
			account.POST("/login", handles.UserLogin)
			account.POST("/register", handles.UserRegister)
			account.POST("/sendcode", handles.SendVerificationCode)
		}
	}
	return r
}
