package models

type GetAccountStatusCredentials struct {
	Username  string `json:"username"`
	Secretkey string `json:"turnstile_secretkey"`
}

type UpdateUsernameCredentials struct {
	Username string `json:"username"`
}
