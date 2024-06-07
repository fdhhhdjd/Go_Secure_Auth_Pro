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

// GenerateOTP generates a random one-time password (OTP) of the specified length.
// It uses the current time as the seed for the random number generator.
// The generated OTP consists of numeric digits only.
func GenerateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
	otp := ""
	for i := 0; i < length; i++ {
		otp += fmt.Sprintf("%d", rand.Intn(10))
	}
	return otp
}

// RandomExpireDuration generates a random expiration duration based on the given number of days.
// It returns a time.Duration representing the random expiration duration.
func RandomExpireDuration(day int) time.Duration {
	days := rand.Intn(day)   // Random ngày trong khoảng day ngày tới
	hours := rand.Intn(24)   // Random giờ trong ngày từ 0 đến 23 giờ
	minutes := rand.Intn(60) // Random phút trong giờ từ 0 đến 59 phút
	seconds := rand.Intn(60) // Random giây trong phút từ 0 đến 59 giây
	expireDuration := time.Duration(days*24)*time.Hour +
		time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second
	return expireDuration
}
