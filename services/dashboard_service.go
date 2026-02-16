package services

import (
	"context"

	"property-backend/models"
	"property-backend/repositories"
)

// DashboardService provides dashboard metrics specific to a user
// It aggregates counts from multiple domains for the authenticated user.
type DashboardService interface {
	TotalPropertyCount(ctx context.Context, userID uint) (int64, error)
	TotalAssetCount(ctx context.Context, userID uint) (int64, error)
	TotalAMCContractCount(ctx context.Context, userID uint) (int64, error)
	TotalActiveRentalCount(ctx context.Context, userID uint) (int64, error)
	RecentActivities(ctx context.Context, userID uint) ([]models.ActivityItem, error)
}

type dashboardService struct {
	dashboardRepo repositories.DashboardRepository
}

// NewDashboardService constructs a DashboardService
func NewDashboardService(dashboardRepo repositories.DashboardRepository) DashboardService {
	return &dashboardService{
		dashboardRepo: dashboardRepo,
	}
}

// TotalPropertyCount returns total property count for the user
func (s *dashboardService) TotalPropertyCount(ctx context.Context, userID uint) (int64, error) {
	return s.dashboardRepo.TotalPropertyCount(ctx, userID)
}

// TotalAssetCount returns total asset count for the user
func (s *dashboardService) TotalAssetCount(ctx context.Context, userID uint) (int64, error) {
	return s.dashboardRepo.TotalAssetCount(ctx, userID)
}

// TotalAMCContractCount returns total AMC contract count for the user
func (s *dashboardService) TotalAMCContractCount(ctx context.Context, userID uint) (int64, error) {
	return s.dashboardRepo.TotalAMCContractCount(ctx, userID)
}

// TotalActiveRentalCount returns total active rental count for the user
func (s *dashboardService) TotalActiveRentalCount(ctx context.Context, userID uint) (int64, error) {
	return s.dashboardRepo.TotalActiveRentalCount(ctx, userID)
}

// RecentActivities returns recent activities for the user
func (s *dashboardService) RecentActivities(ctx context.Context, userID uint) ([]models.ActivityItem, error) {
	const limit = 6
	return s.dashboardRepo.RecentActivities(ctx, userID, limit)
}
