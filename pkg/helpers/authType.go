package helpers

import (
	"net/mail"
	"regexp"
	"strings"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
)

// IdentifyType identifies the type of the given identification string.
// It checks if the identification string is an email, phone number, or username.
// If it's an email, it returns constants.Email.
// If it's a phone number, it returns constants.Phone.
// If it's neither an email nor a phone number, it assumes it's a username and returns constants.Username.
func IdentifyType(identify string) int {
	// Check if it's an email
	_, err := mail.ParseAddress(identify)
	if err == nil {
		return constants.Email
	}

	// Check if it's a phone number
	phoneRegex := `^\+?([0-9]{1,3})?[-. ]?([0-9]{1,4})[-. ]?([0-9]{1,4})[-. ]?([0-9]{9,10})$`
	match, _ := regexp.MatchString(phoneRegex, identify)
	if match {
		return constants.Phone
	}

	// If it's not an email or a phone number, assume it's a username
	return constants.Username
}

// HideEmail hides the username portion of an email address by replacing all but the last two characters with asterisks.
// It takes an email address as input and returns the modified email address as a string.
// If the input email address is not valid (i.e., does not contain an "@" symbol), it returns the original email address.

func HideEmail(email string) string {
	at := strings.LastIndex(email, "@")
	if at == -1 {
		return email // not a valid email
	}

	username := email[:at]
	domain := email[at:]

	if len(username) > 2 {
		username = strings.Repeat("*", len(username)-2) + username[len(username)-2:]
	}

	return username + domain
}
