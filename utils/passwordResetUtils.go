package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// GeneratePasswordResetToken generates a secure random token for password reset
func GeneratePasswordResetToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GetPasswordResetTokenExpiry returns the expiry time for password reset token (24 hours from now)
func GetPasswordResetTokenExpiry() time.Time {
	return time.Now().Add(24 * time.Hour)
}

// BuildPasswordResetEmailTemplate builds the HTML email template for password reset
func BuildPasswordResetEmailTemplate(firstName, resetLink string) string {
	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<style>
				body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
				.container { max-width: 600px; margin: 0 auto; padding: 20px; }
				.header { background-color: #007bff; color: white; padding: 20px; text-align: center; border-radius: 5px; }
				.content { padding: 20px; border: 1px solid #ddd; border-radius: 5px; margin-top: 20px; }
				.button { display: inline-block; background-color: #007bff; color: white; padding: 12px 24px; text-decoration: none; border-radius: 5px; margin-top: 20px; }
				.footer { margin-top: 20px; font-size: 12px; color: #666; }
				.warning { color: #ff6b6b; font-weight: bold; }
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Password Reset Request</h1>
				</div>
				<div class="content">
					<p>Hello %s,</p>
					<p>We received a request to reset the password for your account. If you didn't make this request, you can ignore this email.</p>
					<p>Click the button below to reset your password:</p>
					<a href="%s" class="button">Reset Password</a>
					<p><span class="warning">Note:</span> This link will expire in 24 hours.</p>
					<p>Or copy and paste this link in your browser:</p>
					<p><code>%s</code></p>
				</div>
				<div class="footer">
					<p>This is an automated email. Please do not reply to this address.</p>
					<p>&copy; 2026 Property Management System. All rights reserved.</p>
				</div>
			</div>
		</body>
		</html>
	`, firstName, resetLink, resetLink)
}

// SendPasswordResetEmail sends a password reset email via Twilio SendGrid
func SendPasswordResetEmail(toEmail, firstName, resetToken string) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("SENDGRID_API_KEY not configured")
	}

	fromEmail := os.Getenv("SENDGRID_FROM_EMAIL")
	if fromEmail == "" {
		fromEmail = "noreply@propertymanagement.com"
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", frontendURL, resetToken)
	emailTemplate := BuildPasswordResetEmailTemplate(firstName, resetLink)

	from := mail.NewEmail("Property Management", fromEmail)
	to := mail.NewEmail(firstName, toEmail)
	message := mail.NewSingleEmail(from, "Password Reset Request", to, "Password Reset Request", emailTemplate)
	message.SetReplyTo(mail.NewEmail("Support", fromEmail))

	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		return err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		log.Printf("SendGrid error: Status %d, Body: %s\n", response.StatusCode, response.Body)
		return fmt.Errorf("sendgrid error: status %d", response.StatusCode)
	}

	log.Printf("✅ Password reset email sent successfully to %s\n", toEmail)
	return nil
}
