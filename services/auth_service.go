package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"property-backend/models"
	"property-backend/repositories"
	"property-backend/utils"
)

// AuthResponse contains user info and JWT token
type AuthResponse struct {
	Token     string       `json:"token"`
	User      *models.User `json:"user"`
	ExpiresIn int          `json:"expires_in"`
}

// AuthService defines authentication related business logic
type AuthService interface {
	SignUp(ctx context.Context, u *models.User) (*AuthResponse, error)
	SignIn(ctx context.Context, email, password string) (*AuthResponse, error)
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, newPassword string) error
}

type authService struct {
	repo repositories.AuthRepository
}

// NewAuthService constructs an AuthService
func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) SignUp(ctx context.Context, u *models.User) (*AuthResponse, error) {
	// Create user in database
	id, err := s.repo.Create(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Set the user ID from the created user
	u.UserID = uint(id)

	// Generate JWT token
	token, err := utils.GenerateJWT(u.UserID, u.Email, u.FirstName, u.LastName, u.RoleID)
	if err != nil {
		log.Printf("Warning: JWT generation failed during signup: %v\n", err)
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	// Return response with token and user info (excluding sensitive data)
	return &AuthResponse{
		Token:     token,
		User:      u,
		ExpiresIn: 86400, // 24 hours in seconds
	}, nil
}

func (s *authService) SignIn(ctx context.Context, email, password string) (*AuthResponse, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	// verify password using bcrypt
	if user == nil || !utils.CheckPasswordHash(password, user.HashedPassword) {
		return nil, ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.UserID, user.Email, user.FirstName, user.LastName, user.RoleID)
	if err != nil {
		log.Printf("Warning: JWT generation failed during signin: %v\n", err)
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	// Return response with token and user info
	return &AuthResponse{
		Token:     token,
		User:      user,
		ExpiresIn: 86400, // 24 hours in seconds
	}, nil
}

// ForgotPassword generates a reset token and sends it via email
func (s *authService) ForgotPassword(ctx context.Context, email string) error {
	// Check if user exists
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	if user == nil {
		return fmt.Errorf("user not found")
	}

	// Generate reset token
	token, err := utils.GeneratePasswordResetToken()
	if err != nil {
		return fmt.Errorf("failed to generate reset token")
	}

	// Get expiry time (24 hours from now)
	expiry := utils.GetPasswordResetTokenExpiry()

	// Store token in database
	if err := s.repo.UpdatePasswordResetToken(ctx, email, token, &expiry); err != nil {
		return fmt.Errorf("failed to save reset token")
	}

	// Send email via Gmail SMTP
	if err := utils.SendPasswordResetEmail(email, user.FirstName, token); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

// ResetPassword verifies the token and updates the password
func (s *authService) ResetPassword(ctx context.Context, token, newPassword string) error {
	// Get user by reset token
	user, err := s.repo.GetByResetToken(ctx, token)
	if err != nil {
		return fmt.Errorf("invalid reset token")
	}
	if user == nil {
		return fmt.Errorf("invalid reset token")
	}

	// Check if token is expired
	if user.PasswordResetExpiry == nil || time.Now().After(*user.PasswordResetExpiry) {
		return fmt.Errorf("reset token expired")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password")
	}

	// Update password and clear reset token
	if err := s.repo.UpdatePassword(ctx, user.UserID, hashedPassword); err != nil {
		return fmt.Errorf("failed to update password")
	}

	return nil
}
