package validate

import "regexp"

func IsValidPassword(password string) bool {
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()]`).MatchString(password)

	return hasLowercase && hasUppercase && hasSpecial
}
