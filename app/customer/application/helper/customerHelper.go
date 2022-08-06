package helper

import (
	"crypto/sha256"
	"encoding/base64"
)

func CreatePasswordHash(password string) string {
	passwordByte := []byte(password)
	sha := sha256.New()
	sha.Write(passwordByte)
	passwordHash := base64.URLEncoding.EncodeToString(sha.Sum(nil))
	return passwordHash
}