package repositories

import (
	"context"

	"property-backend/models"

	"gorm.io/gorm"
)

// AssetRepository defines asset-related data access methods
type AssetRepository interface {
	Create(ctx context.Context, a *models.Asset) (int64, error)
	ListAll(ctx context.Context) ([]models.Asset, error)
	CountAll(ctx context.Context) (int64, error)
}

type assetRepository struct {
	db *gorm.DB
}

// NewAssetRepository constructs an AssetRepository
func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) Create(ctx context.Context, a *models.Asset) (int64, error) {
	if err := r.db.WithContext(ctx).Create(a).Error; err != nil {
		return 0, err
	}
	return int64(a.AssetID), nil
}

func (r *assetRepository) ListAll(ctx context.Context) ([]models.Asset, error) {
	var assets []models.Asset
	if err := r.db.WithContext(ctx).Find(&assets).Error; err != nil {
		return nil, err
	}
	return assets, nil
}

func (r *assetRepository) CountAll(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Asset{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
