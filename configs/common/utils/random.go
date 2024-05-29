package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := base64.StdEncoding.EncodeToString(b)
	return token, nil
}
