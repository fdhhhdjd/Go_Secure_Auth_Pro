package helpers

import (
	"math/rand"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"golang.org/x/crypto/bcrypt"
)

var charset = global.Cfg.Server.KeyPassword

// generatePassword generates a random password of the specified length.
func generatePassword(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateRandomPassword generates a random password with the specified length.
func GenerateRandomPassword(character int) string {
	return generatePassword(character)
}

// HashPassword generates a salted and hashed password using bcrypt.
// It takes the plain-text password and the number of salt rounds as input.
// It returns the generated salt, the hashed password, and any error encountered.
func HashPassword(password string, saltRounds int) (string, string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(password), saltRounds)
	if err != nil {
		return "", "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return string(salt), string(hashedPassword), nil
}

// ComparePassword compares a plain-text password with a hashed password and returns true if they match.
// It uses bcrypt.CompareHashAndPassword to perform the comparison.
func ComparePassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

// HashedPasswordOld generates a hashed password using bcrypt with the provided password and salt.
// It returns the hashed password as a string and any error encountered during the hashing process.
func HashedPasswordOld(password string, salt string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
