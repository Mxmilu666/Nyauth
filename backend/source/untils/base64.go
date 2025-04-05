package untils

import (
	"encoding/base64"
	"strings"
)

// base64URLEncode 进行Base64URL编码（RFC7515）
func Base64URLEncode(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	// 替换标准Base64中不符合URL安全的字符
	encoded = strings.ReplaceAll(encoded, "+", "-")
	encoded = strings.ReplaceAll(encoded, "/", "_")
	encoded = strings.TrimRight(encoded, "=")
	return encoded
}
