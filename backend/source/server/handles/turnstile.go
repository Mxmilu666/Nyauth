package handles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"nyauth_backed/source"
)

const turnstileVerifyURL = "https://challenges.cloudflare.com/turnstile/v0/siteverify"

type TurnstileResponse struct {
	Success bool `json:"success"`
}

func VerifyTurnstile(token string) (bool, error) {
	secretKey := source.AppConfig.Turnstile.SecretKey
	if secretKey == "" {
		return false, fmt.Errorf("missing TURNSTILE_SECRET_KEY environment variable")
	}

	reqBody, err := json.Marshal(map[string]string{
		"secret":   secretKey,
		"response": token,
	})
	if err != nil {
		return false, err
	}

	resp, err := http.Post(turnstileVerifyURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var turnstileResp TurnstileResponse
	if err := json.NewDecoder(resp.Body).Decode(&turnstileResp); err != nil {
		return false, err
	}

	return turnstileResp.Success, nil
}
