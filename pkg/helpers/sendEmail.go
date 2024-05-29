package helpers

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
)

// SendEmail sends an email using the configured SMTP settings.
func SendEmail(email string, data string) {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		global.Cfg.Gmail.Mail, // Pass the email address as a string
		global.Cfg.Gmail.Password,
		global.Cfg.Gmail.Host,
	)

	to := []string{email}
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Hello!\r\n" +
		"\r\n" +
		"This is content:\r\n" +
		"Data: " + data + "\r\n")
	err := smtp.SendMail(fmt.Sprintf("%s:%s", global.Cfg.Gmail.Host, global.Cfg.Gmail.Port), auth, "profile-forme.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
