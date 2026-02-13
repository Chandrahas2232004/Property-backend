package services

import (
	"context"
	"sort"

	"property-backend/models"
	"property-backend/repositories"
)

// DashboardService provides dashboard metrics
// It aggregates counts from multiple domains.
type DashboardService interface {
	TotalPropertyCount(ctx context.Context) (int64, error)
	TotalAssetCount(ctx context.Context) (int64, error)
	TotalAMCContractCount(ctx context.Context) (int64, error)
	TotalActiveRentalCount(ctx context.Context) (int64, error)
	RecentActivities(ctx context.Context, userID uint) ([]models.ActivityItem, error)
}

type dashboardService struct {
	propertyRepo  repositories.PropertyRepository
	assetRepo     repositories.AssetRepository
	contractRepo  repositories.ContractRepository
	agreementRepo repositories.AgreementRepository
}

// NewDashboardService constructs a DashboardService
func NewDashboardService(
	propertyRepo repositories.PropertyRepository,
	assetRepo repositories.AssetRepository,
	contractRepo repositories.ContractRepository,
	agreementRepo repositories.AgreementRepository,
) DashboardService {
	return &dashboardService{
		propertyRepo:  propertyRepo,
		assetRepo:     assetRepo,
		contractRepo:  contractRepo,
		agreementRepo: agreementRepo,
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

func (s *dashboardService) RecentActivities(ctx context.Context, userID uint) ([]models.ActivityItem, error) {
	const limit = 6

	propertyItems, err := s.propertyRepo.RecentByUser(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	assetItems, err := s.assetRepo.RecentByUser(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	contractItems, err := s.contractRepo.RecentByUser(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	agreementItems, err := s.agreementRepo.RecentByUser(ctx, userID, limit)
	if err != nil {
		return nil, err
	}

	items := make([]models.ActivityItem, 0, len(propertyItems)+len(assetItems)+len(contractItems)+len(agreementItems))
	items = append(items, propertyItems...)
	items = append(items, assetItems...)
	items = append(items, contractItems...)
	items = append(items, agreementItems...)

	sort.Slice(items, func(i, j int) bool {
		return items[i].TimeCreated.After(items[j].TimeCreated)
	})

	if len(items) > limit {
		items = items[:limit]
	}

	return items, nil
}
