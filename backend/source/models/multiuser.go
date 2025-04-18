package models

// MultiuserCredentials 多用户身份创建参数
type MultiuserCredentials struct {
	DisplayName string `json:"display_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Description string `json:"description"`
	// Avatar      string `json:"avatar"`
	Code string `json:"code" binding:"required"` // 邮箱验证码
}
