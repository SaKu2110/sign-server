package service

import(
	"encoding/hex"
	"crypto/sha256"
)

func CreateHashWithPassord(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
