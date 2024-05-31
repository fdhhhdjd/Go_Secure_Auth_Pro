package utils

import (
	"math/rand"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
)

func GenerateToken() (string, error) {
	const letters = constants.KeyRandomTokenVerification
	b := make([]byte, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	token := string(b)
	return token, nil
}
