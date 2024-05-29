package pkg

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"strconv"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/go-gomail/gomail"
)

// EmailData holds the data for the email template
// SendGoEmail sends an email using the provided email address and email data.
// It ensures that the templates directory exists and parses the HTML template from the string.
// Then, it executes the template with the provided data and initializes the email message.
// The sender, recipient, subject, and body of the email are set accordingly.
// It also embeds an image and sets up the SMTP connection.
// Finally, it sends the email using the SMTP connection.
func SendGoEmail(email string, data models.EmailData) {
	// Ensure the templates directory exists (if you still need to create any directories or files)
	err := os.MkdirAll("templates", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the HTML template from the string
	tmpl, err := template.New("email").Parse(data.Template)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the template with data
	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the email message
	m := gomail.NewMessage()

	// Set the sender
	m.SetHeader("From", global.Cfg.Gmail.Mail)

	// Set the recipient
	m.SetHeader("To", email)

	// Set the subject
	m.SetHeader("Subject", data.Title+"!")

	// Set the body with HTML content
	m.SetBody("text/html", body.String())

	// Embed the image
	m.Embed("docs/assets/logo.png", gomail.SetHeader(map[string][]string{
		"Content-ID": {"<logo>"},
	}))

	// Set up the SMTP connection
	port, err := strconv.Atoi(global.Cfg.Gmail.Port)
	if err != nil {
		panic(err)
	}

	d := gomail.NewDialer(global.Cfg.Gmail.Host, port, global.Cfg.Gmail.Mail, global.Cfg.Gmail.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
