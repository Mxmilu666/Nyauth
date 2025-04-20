package models

// TOTPVerifyRequest 验证TOTP请求
type TOTPVerifyRequest struct {
	Code string `json:"code" binding:"required"`
}

// TOTPLoginRequest TOTP登录请求
type TOTPLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

// TOTPDisableRequest 禁用TOTP请求
type TOTPDisableRequest struct {
	Code string `json:"code" binding:"required"`
}

// TOTPStatus TOTP状态
type TOTPStatus struct {
	Enabled bool `json:"enabled"`
}
