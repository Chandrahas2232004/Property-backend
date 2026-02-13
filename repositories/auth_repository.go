package repositories

import (
	"context"
	"time"

	"property-backend/models"

	"gorm.io/gorm"
)

// AuthRepository defines auth/user-related data access methods
type AuthRepository interface {
	Create(ctx context.Context, u *models.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	UpdatePasswordResetToken(ctx context.Context, email, token string, expiry *time.Time) error
	GetByResetToken(ctx context.Context, token string) (*models.User, error)
	UpdatePassword(ctx context.Context, userID uint, hashedPassword string) error
}

type authRepository struct {
	db *gorm.DB
}

// NewAuthRepository constructs an AuthRepository
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Create(ctx context.Context, u *models.User) (int64, error) {
	if err := r.db.WithContext(ctx).Create(u).Error; err != nil {
		return 0, err
	}
	return int64(u.UserID), nil
}

func (r *authRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdatePasswordResetToken updates the reset token and expiry for a user
func (r *authRepository) UpdatePasswordResetToken(ctx context.Context, email, token string, expiry *time.Time) error {
	return r.db.WithContext(ctx).Model(&models.User{}).
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"password_reset_token":  token,
			"password_reset_expiry": expiry,
		}).Error
}

// GetByResetToken retrieves a user by their reset token
func (r *authRepository) GetByResetToken(ctx context.Context, token string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("password_reset_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdatePassword updates the hashed password for a user and clears the reset token
func (r *authRepository) UpdatePassword(ctx context.Context, userID uint, hashedPassword string) error {
	return r.db.WithContext(ctx).Model(&models.User{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"hashed_password":       hashedPassword,
			"password_reset_token":  nil,
			"password_reset_expiry": nil,
		}).Error
}
