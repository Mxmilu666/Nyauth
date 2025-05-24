package models

type LoginCredentials struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	TotpCode  string `json:"totp_code,omitempty"`
	Secretkey string `json:"turnstile_secretkey"`
}

type EmailCredentials struct {
	Useremail string `json:"useremail"`
	Secretkey string `json:"turnstile_secretkey"`
}

type RegisterCredentials struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Useremail string `json:"useremail"`
	TempCode  string `json:"code"`
	Secretkey string `json:"turnstile_secretkey"`
}

type VerifyCodeCredentials struct {
	Useremail string `json:"useremail" binding:"required,email"`
	Code      string `json:"code" binding:"required"`
}
