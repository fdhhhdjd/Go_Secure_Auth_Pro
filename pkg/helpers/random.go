package helpers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
)

// GenerateToken generates a random token of length 32 using the characters specified in the constants.KeyRandomTokenVerification.
// It returns the generated token as a string and any error encountered during the generation process.
func GenerateToken() (string, error) {
	const letters = constants.KeyRandomTokenVerification
	b := make([]byte, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	token := string(b)
	return token, nil
}

// RandomEmail generates a random email.
// It uses the current time as the seed for the random number generator.
// Returns the generated email as a string.
func RandomEmail() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randPart := r.Intn(100000)
	email := fmt.Sprintf("user%d@example.com", randPart)
	return email
}

// RandomPassword generates a random password.
// It uses the current time as the seed for the random number generator.
// Returns the generated password as a string.
func RandomPassword() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randPart := r.Intn(100000)
	password := fmt.Sprintf("password%d", randPart)
	return password
}
