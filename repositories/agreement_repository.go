package repositories

import (
	"context"

	"gorm.io/gorm"
	"property-backend/models"
)

// AgreementRepository defines agreement-related data access methods
type AgreementRepository interface {
	CreateRental(ctx context.Context, a *models.Agreement) (int64, error)
	ListAll(ctx context.Context) ([]models.Agreement, error)
}

type agreementRepository struct {
	db *gorm.DB
}

// NewAgreementRepository constructs an AgreementRepository
func NewAgreementRepository(db *gorm.DB) AgreementRepository {
	return &agreementRepository{db: db}
}

func (r *agreementRepository) CreateRental(ctx context.Context, a *models.Agreement) (int64, error) {
	if err := r.db.WithContext(ctx).Create(a).Error; err != nil {
		return 0, err
	}
	return int64(a.AgreementID), nil
}

func (r *agreementRepository) ListAll(ctx context.Context) ([]models.Agreement, error) {
	var agreements []models.Agreement
	if err := r.db.WithContext(ctx).Preload("Property").Find(&agreements).Error; err != nil {
		return nil, err
	}
	return agreements, nil
}
