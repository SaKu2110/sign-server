package service

import(
	"crypto/sha256"
)

func CreateHashWithPassord(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}
