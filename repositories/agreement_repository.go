package repositories

import (
	"context"

	"property-backend/models"

	"gorm.io/gorm"
)

// AgreementRepository defines agreement-related data access methods
type AgreementRepository interface {
	CreateRental(ctx context.Context, a *models.Agreement) (int64, error)
	ListAll(ctx context.Context) ([]models.Agreement, error)
	RecentByUser(ctx context.Context, userID uint, limit int) ([]models.ActivityItem, error)
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

func (r *agreementRepository) RecentByUser(ctx context.Context, userID uint, limit int) ([]models.ActivityItem, error) {
	var items []models.ActivityItem
	if err := r.db.WithContext(ctx).
		Table("agreement").
		Select("agreement.agreement_id as respective_id, agreement.created_at as time_created, ? as type", "agreement").
		Joins("JOIN property ON property.property_id = agreement.property_id").
		Where("property.user_id = ?", userID).
		Order("agreement.created_at desc").
		Limit(limit).
		Scan(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
