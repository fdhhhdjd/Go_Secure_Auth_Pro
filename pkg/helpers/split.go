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

// HidePhoneNumber hides the first six digits of the given phone number and only shows the last four digits.
// If the phone number is less than 10 digits, it returns the original phone number.
func HidePhoneNumber(phone string) string {
	if len(phone) < 10 {
		return phone
	}
	return strings.Repeat("*", len(phone)-4) + phone[len(phone)-4:]
}
