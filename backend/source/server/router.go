package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(filterLogs())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	api := r.Group("/api/v0")
	{
		api.GET("/code", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello, World!")
		})

		account := api.Group("/account")
		{
			account.POST("/login", func(c *gin.Context) {
				c.String(http.StatusOK, "Hello, World!")
			})
		}
	}
	return r
}
