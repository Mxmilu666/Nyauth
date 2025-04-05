package oauth

import (
	"nyauth_backed/source/untils"
	"strings"
	"sync"
	"time"
)

// Token 结构体用于存储访问令牌信息
type Token struct {
	AccessToken string    // 访问令牌
	ClientID    string    // 客户端ID
	UserID      string    // 用户ID
	Scope       []string  // 权限范围，修改为字符串数组
	Exp         time.Time // 过期时间
}

var (
	// 内存存储令牌
	tokens     = make(map[string]*Token)
	tokenMutex sync.RWMutex
)

// CreateToken 创建新的访问令牌并存储在内存中
func CreateToken(clientID, userID string, scope []string, expiresIn ...int) (string, error) {
	token, err := untils.GenerateRandomCode(48, false) // 生成随机字符串作为访问令牌
	if err != nil {
		return "", err
	}

	// 默认过期时间为2小时
	exp := 7200
	if len(expiresIn) > 0 && expiresIn[0] > 0 {
		exp = expiresIn[0]
	}

	// 创建令牌对象
	accessToken := &Token{
		AccessToken: token,
		ClientID:    clientID,
		UserID:      userID,
		Scope:       scope,
		Exp:         time.Now().Add(time.Duration(exp) * time.Second),
	}

	// 存储到内存中
	tokenMutex.Lock()
	tokens[token] = accessToken
	tokenMutex.Unlock()

	return token, nil
}

// 为了兼容性，提供一个接收逗号分隔字符串的创建方法
func CreateTokenFromString(clientID, userID, scopeStr string, expiresIn ...int) (string, error) {
	var scope []string
	if scopeStr != "" {
		scope = strings.Split(scopeStr, ",")
	} else {
		scope = []string{}
	}
	return CreateToken(clientID, userID, scope, expiresIn...)
}

// GetToken 通过访问令牌获取对应的信息
func GetToken(token string) (*Token, bool) {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()

	accessToken, exists := tokens[token]
	if !exists || time.Now().After(accessToken.Exp) {
		// 如果不存在或已过期
		return nil, false
	}

	return accessToken, true
}

// HasScope 检查令牌是否具有指定的权限范围
func (t *Token) HasScope(scope string) bool {
	for _, s := range t.Scope {
		if s == scope {
			return true
		}
	}
	return false
}

// GetScopeString 获取逗号分隔的权限范围字符串
func (t *Token) GetScopeString() string {
	return strings.Join(t.Scope, ",")
}

// RemoveToken 移除访问令牌
func RemoveToken(token string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	delete(tokens, token)
}

// 定期清理过期的访问令牌
func init() {
	go func() {
		for {
			time.Sleep(15 * time.Minute) // 每15分钟清理一次
			cleanupExpiredTokens()
		}
	}()
}

// cleanupExpiredTokens 清理过期的访问令牌
func cleanupExpiredTokens() {
	now := time.Now()
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	for tokenStr, token := range tokens {
		if now.After(token.Exp) {
			delete(tokens, tokenStr)
		}
	}
}
