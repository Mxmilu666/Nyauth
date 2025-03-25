package untils

import (
	"crypto/sha1"

	"github.com/google/uuid"
)

func ToUUIDv5(id string) string {
	// 命名空间 UUID
	// 有私货 (?
	namespace := uuid.MustParse("72796f79-616d-479a-8646-1c6f76650000")
	hash := sha1.New()
	hash.Write(namespace[:])
	hash.Write([]byte(id))
	hashBytes := hash.Sum(nil)

	var u uuid.UUID
	copy(u[:], hashBytes[:16])

	u[6] = (u[6] & 0x0F) | 0x50
	u[8] = (u[8] & 0x3F) | 0x80

	return u.String()
}
