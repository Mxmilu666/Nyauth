package untils

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

// GenerateRandomCode 生成随机字符串
func GenerateRandomCode(length int, toUpper bool) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// 先使用base64编码
	encodedStr := base64.URLEncoding.EncodeToString(bytes)[:length]

	// 根据参数决定是否转换为大写
	if toUpper {
		return strings.ToUpper(encodedStr), nil
	}
	return encodedStr, nil
}
