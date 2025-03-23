package untils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(strings.ToLower(text)))
	return hex.EncodeToString(hasher.Sum(nil))
}
