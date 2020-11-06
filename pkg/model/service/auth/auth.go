package auth

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/SaKu2110/sign-server/pkg/model/dto"
)

func CheckPassword(hash dto.UserInfo, rawPassword string) bool {
	return hash.Password == Hash(rawPassword)
}

func Hash(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}
