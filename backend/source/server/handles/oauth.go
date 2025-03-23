package handles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OAuthAuthorize(c *gin.Context) {
	SendResponse(c, http.StatusBadRequest, "你说得对但是后面忘了", nil)
}
