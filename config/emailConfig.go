package config

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// EmailConfig holds Gmail SMTP configuration
type EmailConfig struct {
	SMTPServer   string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	FromEmail    string
}

var GmailConfig *EmailConfig

// InitGmail initializes the Gmail SMTP configuration
func InitGmail() {
	smtpUsername := os.Getenv("GMAIL_USERNAME")
	smtpPassword := os.Getenv("GMAIL_PASSWORD")

	if smtpUsername == "" || smtpPassword == "" {
		log.Println("⚠️  GMAIL_USERNAME or GMAIL_PASSWORD not set - email functionality will be disabled")
		return
	}

	GmailConfig = &EmailConfig{
		SMTPServer:   "smtp.gmail.com",
		SMTPPort:     "587",
		SMTPUsername: smtpUsername,
		SMTPPassword: smtpPassword,
		FromEmail:    smtpUsername,
	}
	log.Println("✅ Gmail SMTP initialized successfully")
}

// SendEmail sends an email using Gmail SMTP
func SendEmail(fromEmail, toEmail, subject, htmlContent string) error {
	if GmailConfig == nil {
		return fmt.Errorf("Gmail configuration not initialized")
	}

	auth := smtp.PlainAuth("", GmailConfig.SMTPUsername, GmailConfig.SMTPPassword, GmailConfig.SMTPServer)
	addr := fmt.Sprintf("%s:%s", GmailConfig.SMTPServer, GmailConfig.SMTPPort)

	headers := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n",
		GmailConfig.FromEmail, toEmail, subject)
	body := headers + htmlContent

	err := smtp.SendMail(addr, auth, GmailConfig.FromEmail, []string{toEmail}, []byte(body))
	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	log.Printf("✅ Email sent successfully to %s\n", toEmail)
	return nil
}
