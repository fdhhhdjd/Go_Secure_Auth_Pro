package validate

import (
	"regexp"
)

// ValidateAndRespond validates a value using the provided validation function and returns a boolean indicating whether the value is valid.
// If the value is not empty and fails the validation, it returns false. Otherwise, it returns true.
func ValidateAndRespond(value string, validateFunc func(string) bool) bool {
	if value != "" {
		if !validateFunc(value) {
			return false
		}
	}
	return true
}

// IsValidPassword checks if a password meets the required criteria.
// It returns true if the password contains at least one lowercase letter,
// one uppercase letter, and one special character, otherwise it returns false.
func IsValidPassword(password string) bool {
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()]`).MatchString(password)

	return hasLowercase && hasUppercase && hasSpecial
}

// IsValidateUser checks if a username is valid.
// A valid username is 3-16 characters long and contains only alphanumeric characters.
func IsValidateUser(username string) bool {
	// A valid username is 3-16 characters long and contains only alphanumeric characters
	isValid := regexp.MustCompile(`^[a-zA-Z0-9]{3,16}$`).MatchString(username)
	return isValid
}

// IsValidatePhone checks if a given phone number is a valid Vietnamese phone number.
// A valid Vietnamese phone number starts with 0 and is 10 or 11 digits long.
func IsValidatePhone(phone string) bool {
	// A valid Vietnamese phone number starts with 0 and is 10 or 11 digits long
	isValid := regexp.MustCompile(`^0[0-9]{9,10}$`).MatchString(phone)
	return isValid
}
