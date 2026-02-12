package repositories

import (
	"context"

	"property-backend/models"

	"gorm.io/gorm"
)

// ContractRepository defines contract-related data access methods
type ContractRepository interface {
	Create(ctx context.Context, c *models.Contract) (int64, error)
	ListAll(ctx context.Context) ([]models.Contract, error)
	ListByType(ctx context.Context, contractType string) ([]models.Contract, error)
	CountByType(ctx context.Context, contractType string) (int64, error)
}

type contractRepository struct {
	db *gorm.DB
}

// NewContractRepository constructs a ContractRepository
func NewContractRepository(db *gorm.DB) ContractRepository {
	return &contractRepository{db: db}
}

func (r *contractRepository) Create(ctx context.Context, c *models.Contract) (int64, error) {
	if err := r.db.WithContext(ctx).Create(c).Error; err != nil {
		return 0, err
	}
	return int64(c.ContractID), nil
}

func (r *contractRepository) ListAll(ctx context.Context) ([]models.Contract, error) {
	var contracts []models.Contract
	if err := r.db.WithContext(ctx).Preload("Asset").Preload("ContractType").Find(&contracts).Error; err != nil {
		return nil, err
	}
	return contracts, nil
}

func (r *contractRepository) ListByType(ctx context.Context, contractType string) ([]models.Contract, error) {
	var contracts []models.Contract
	if err := r.db.WithContext(ctx).
		Preload("Asset").
		Preload("ContractType").
		Joins("ContractType").
		Where("contract_type_master.contract_type_name = ?", contractType).
		Find(&contracts).Error; err != nil {
		return nil, err
	}
	return contracts, nil
}

func (r *contractRepository) CountByType(ctx context.Context, contractType string) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&models.Contract{}).
		Joins("LEFT JOIN contract_type_master ON contract_type_master.contract_type_id = contracts.contract_type_id").
		Where("contract_type_master.contract_type_name = ?", contractType).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
