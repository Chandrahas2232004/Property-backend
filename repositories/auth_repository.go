package repositories

import (
	"context"

	"gorm.io/gorm"
	"property-backend/models"
)

// AuthRepository defines auth/user-related data access methods
type AuthRepository interface {
	Create(ctx context.Context, u *models.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
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
