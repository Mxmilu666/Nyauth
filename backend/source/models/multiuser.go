package models

// 解析多用户请求体中的数据
var MultiuserCredentials struct {
	DisplayName string `json:"display_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Description string `jsson:"description"`
	Avatar      string `json:"avatar"`
}
