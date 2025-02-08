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
