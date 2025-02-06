package handles

import (
	"encoding/json"
	"net/http"
	"nyauth_backed/source/database"
)

type Credentials struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Secretkey string `json:"turnstile_secretkey"`
}

func Userlogin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 看看用户在不在
	userExists := database.CheckUserExists(creds.Username)
	if !userExists {
		http.Error(w, "User does not exist", http.StatusUnauthorized)
		return
	}
}
