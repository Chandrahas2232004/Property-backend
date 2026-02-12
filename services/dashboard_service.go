package services

import (
	"context"

	"property-backend/repositories"
)

// DashboardService provides dashboard metrics
// It aggregates counts from multiple domains.
type DashboardService interface {
	TotalPropertyCount(ctx context.Context) (int64, error)
	TotalAssetCount(ctx context.Context) (int64, error)
	TotalAMCContractCount(ctx context.Context) (int64, error)
	TotalActiveRentalCount(ctx context.Context) (int64, error)
}

type dashboardService struct {
	propertyRepo repositories.PropertyRepository
	assetRepo    repositories.AssetRepository
	contractRepo repositories.ContractRepository
}

// NewDashboardService constructs a DashboardService
func NewDashboardService(
	propertyRepo repositories.PropertyRepository,
	assetRepo repositories.AssetRepository,
	contractRepo repositories.ContractRepository,
) DashboardService {
	return &dashboardService{
		propertyRepo: propertyRepo,
		assetRepo:    assetRepo,
		contractRepo: contractRepo,
	}
}

func (s *dashboardService) TotalPropertyCount(ctx context.Context) (int64, error) {
	return s.propertyRepo.Total(ctx)
}

func (s *dashboardService) TotalAssetCount(ctx context.Context) (int64, error) {
	return s.assetRepo.CountAll(ctx)
}

func (s *dashboardService) TotalAMCContractCount(ctx context.Context) (int64, error) {
	return s.contractRepo.CountByType(ctx, "amc")
}

func (s *dashboardService) TotalActiveRentalCount(ctx context.Context) (int64, error) {
	return s.propertyRepo.ActiveRentalCount(ctx)
}
