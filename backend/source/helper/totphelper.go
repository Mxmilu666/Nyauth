package helper

import (
	"sync"
	"time"
)

// TOTP临时密钥存储
type tempTOTPSecret struct {
	Secret    string
	ExpiresAt time.Time
}

var (
	// 内存缓存，用于存储临时TOTP密钥
	totpCache = struct {
		sync.RWMutex
		m map[string]tempTOTPSecret
	}{m: make(map[string]tempTOTPSecret)}
)

func init() {
	// 启动定期清理过期的TOTP密钥
	go cleanupExpiredTOTPSecrets()
}

// 定期清理过期的TOTP密钥
func cleanupExpiredTOTPSecrets() {
	ticker := time.NewTicker(5 * time.Minute) // 每5分钟清理一次
	defer ticker.Stop()

	for range ticker.C {
		removeExpiredTOTPSecrets()
	}
}

// 清理过期的TOTP密钥
func removeExpiredTOTPSecrets() {
	now := time.Now()

	totpCache.Lock()
	defer totpCache.Unlock()

	for key, secret := range totpCache.m {
		if now.After(secret.ExpiresAt) {
			delete(totpCache.m, key)
		}
	}
}

// SetTempTOTPSecret 设置临时TOTP密钥
func SetTempTOTPSecret(key, secret string, expiresIn time.Duration) error {
	totpCache.Lock()
	defer totpCache.Unlock()

	totpCache.m[key] = tempTOTPSecret{
		Secret:    secret,
		ExpiresAt: time.Now().Add(expiresIn),
	}
	return nil
}

// GetTempTOTPSecret 获取临时TOTP密钥
func GetTempTOTPSecret(key string) (string, bool) {
	totpCache.RLock()
	defer totpCache.RUnlock()

	secret, exists := totpCache.m[key]
	if !exists || time.Now().After(secret.ExpiresAt) {
		return "", false
	}
	return secret.Secret, true
}

// RemoveTempTOTPSecret 删除临时TOTP密钥
func RemoveTempTOTPSecret(key string) {
	totpCache.Lock()
	defer totpCache.Unlock()
	delete(totpCache.m, key)
}
