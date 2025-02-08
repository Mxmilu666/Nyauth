package handles

import "github.com/gin-gonic/gin"

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
