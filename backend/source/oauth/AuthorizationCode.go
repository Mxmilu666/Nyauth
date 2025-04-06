package oauth

import (
	"nyauth_backed/source/untils"
	"sync"
	"time"
)

// AuthorizationCode 结构体用于存储授权码信息
type AuthorizationCode struct {
	Code     string    // 授权码
	ClientID string    // 客户端ID
	UserID   string    // 用户ID
	Exp      time.Time // 过期时间
}

var (
	// 内存存储授权码
	authCodes     = make(map[string]*AuthorizationCode)
	authCodeMutex sync.RWMutex
)

// CreateAuthorizationCode 创建新的授权码并存储在内存中
func CreateAuthorizationCode(clientID, userID string, expiresIn ...int) (string, error) {
	code, err := untils.GenerateRandomCode(32, false) // 生成随机字符串作为授权码
	if err != nil {
		return "", err
	}

	// 默认过期时间为10分钟
	exp := 600
	if len(expiresIn) > 0 && expiresIn[0] > 0 {
		exp = expiresIn[0]
	}

	// 创建授权码对象
	authCode := &AuthorizationCode{
		Code:     code,
		ClientID: clientID,
		UserID:   userID,
		Exp:      time.Now().Add(time.Duration(exp) * time.Second),
	}

	// 存储到内存中
	authCodeMutex.Lock()
	authCodes[code] = authCode
	authCodeMutex.Unlock()

	return code, nil
}

// GetAuthorizationCode 通过授权码获取对应的信息
func GetAuthorizationCode(code string) (*AuthorizationCode, bool) {
	authCodeMutex.RLock()
	defer authCodeMutex.RUnlock()

	authCode, exists := authCodes[code]
	if !exists || time.Now().After(authCode.Exp) {
		// 如果不存在或已过期
		return nil, false
	}

	return authCode, true
}

// RemoveAuthorizationCode 移除授权码
func RemoveAuthorizationCode(code string) {
	authCodeMutex.Lock()
	defer authCodeMutex.Unlock()
	delete(authCodes, code)
}

// 定期清理过期的授权码
func init() {
	go periodicCleanup(cleanupExpiredCodes, 5*time.Minute)
}

// cleanupExpiredCodes 清理过期的授权码
func cleanupExpiredCodes() {
	now := time.Now()
	authCodeMutex.Lock()
	defer authCodeMutex.Unlock()

	for code, authCode := range authCodes {
		if now.After(authCode.Exp) {
			delete(authCodes, code)
		}
	}
}
