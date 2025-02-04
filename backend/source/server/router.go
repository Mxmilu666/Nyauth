package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setuprouter() *gin.Engine {
	r := gin.New()
	r.Use(filterLogs())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	return r
}
