package helpers

import "strings"

// GetUsernameFromEmail extracts the username from the given email address.
// It searches for the first occurrence of "@" in the email string and returns
// the substring before "@" as the username. If no "@" is found, it returns an empty string.
func GetUsernameFromEmail(email string) string {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return ""
	}
	return email[:atIndex]
}
