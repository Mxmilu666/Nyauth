package oauth

import (
	"strings"
	"time"
)

// periodicCleanup 定期执行清理函数
func periodicCleanup(cleanupFunc func(), interval time.Duration) {
	for {
		time.Sleep(interval)
		cleanupFunc()
	}
}

// ValidateScope 验证权限范围
func ValidateScope(userPermissions []string, targetPermission string) bool {
	for _, perm := range userPermissions {
		// 精确匹配
		if perm == targetPermission {
			return true
		}

		// 通配符匹配
		if strings.HasSuffix(perm, ":*") {
			prefix := strings.TrimSuffix(perm, ":*")
			if strings.HasPrefix(targetPermission, prefix+":") {
				return true
			}
		}
	}
	return false
}
