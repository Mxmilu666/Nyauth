package models

type LoginCredentials struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
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
	Code      string `json:"code"`
	Secretkey string `json:"turnstile_secretkey"`
}
