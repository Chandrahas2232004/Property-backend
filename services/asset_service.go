package services

import (
	"context"

	"property-backend/models"
	"property-backend/repositories"
)

// AssetService defines asset domain logic
type AssetService interface {
	AddAsset(ctx context.Context, a *models.Asset) (int64, error)
	GetAllAssets(ctx context.Context) ([]models.Asset, error)
}

type assetService struct {
	repo repositories.AssetRepository
}

// NewAssetService constructs an AssetService
func NewAssetService(repo repositories.AssetRepository) AssetService {
	return &assetService{repo: repo}
}

func (s *assetService) AddAsset(ctx context.Context, a *models.Asset) (int64, error) {
	return s.repo.Create(ctx, a)
}

func (s *assetService) GetAllAssets(ctx context.Context) ([]models.Asset, error) {
	return s.repo.ListAll(ctx)
}
