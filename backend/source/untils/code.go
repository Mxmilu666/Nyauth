package untils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// GenerateRandomCode 生成随机字符串
func GenerateRandomCode(length int, toUpper bool) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

	charsetLength := big.NewInt(int64(len(charset)))
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		result[i] = charset[randomIndex.Int64()]
	}

	resultStr := string(result)
	if toUpper {
		return strings.ToUpper(resultStr), nil
	}
	return resultStr, nil
}
