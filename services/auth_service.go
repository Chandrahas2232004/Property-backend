package services

import (
	"context"

	"property-backend/models"
	"property-backend/repositories"
	"property-backend/utils"
)

// AuthService defines authentication related business logic
type AuthService interface {
	SignUp(ctx context.Context, u *models.User) (int64, error)
	SignIn(ctx context.Context, email, password string) (*models.User, error)
}

type authService struct {
	repo repositories.AuthRepository
}

// NewAuthService constructs an AuthService
func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) SignUp(ctx context.Context, u *models.User) (int64, error) {
	return s.repo.Create(ctx, u)
}

func (s *authService) SignIn(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	// verify password using bcrypt
	if user == nil || !utils.CheckPasswordHash(password, user.HashedPassword) {
		return nil, ErrInvalidCredentials
	}
	return user, nil
}
