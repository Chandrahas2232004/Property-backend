package config

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var SendGridClient *sendgrid.Client

// InitSendGrid initializes the Twilio SendGrid client
func InitSendGrid() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		log.Println("⚠️  SENDGRID_API_KEY not set - email functionality will be disabled")
		return
	}

	SendGridClient = sendgrid.NewSendClient(apiKey)
	log.Println("✅ SendGrid initialized successfully")
}

// SendEmail sends an email using Twilio SendGrid
func SendEmail(fromEmail, toEmail, subject, htmlContent string) error {
	if SendGridClient == nil {
		return fmt.Errorf("SendGrid client not initialized")
	}

	from := mail.NewEmail("Property Management", fromEmail)
	to := mail.NewEmail("User", toEmail)
	message := mail.NewSingleEmail(from, subject, to, subject, htmlContent)

	response, err := SendGridClient.Send(message)
	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		log.Printf("SendGrid error: Status %d, Body: %s\n", response.StatusCode, response.Body)
		return fmt.Errorf("sendgrid error: status %d", response.StatusCode)
	}

	log.Printf("✅ Email sent successfully to %s (Status: %d)\n", toEmail, response.StatusCode)
	return nil
}
