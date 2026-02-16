package repositories

import (
	"context"
	"time"

	"property-backend/models"

	"gorm.io/gorm"
)

// DashboardRepository defines dashboard-related data access methods
type DashboardRepository interface {
	TotalPropertyCount(ctx context.Context, userID uint) (int64, error)
	TotalAssetCount(ctx context.Context, userID uint) (int64, error)
	TotalAMCContractCount(ctx context.Context, userID uint) (int64, error)
	TotalActiveRentalCount(ctx context.Context, userID uint) (int64, error)
	RecentActivities(ctx context.Context, userID uint, limit int) ([]models.ActivityItem, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

// NewDashboardRepository constructs a DashboardRepository
func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

// TotalPropertyCount returns total property count for a specific user
func (r *dashboardRepository) TotalPropertyCount(ctx context.Context, userID uint) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&models.Property{}).
		Where("user_id = ?", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// TotalAssetCount returns total asset count for a specific user
func (r *dashboardRepository) TotalAssetCount(ctx context.Context, userID uint) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&models.Asset{}).
		Joins("JOIN property ON property.property_id = assets.property_id").
		Where("property.user_id = ?", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// TotalAMCContractCount returns total AMC contract count for a specific user
func (r *dashboardRepository) TotalAMCContractCount(ctx context.Context, userID uint) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&models.Contract{}).
		Joins("JOIN assets ON assets.asset_id = contracts.asset_id").
		Joins("JOIN property ON property.property_id = assets.property_id").
		Joins("LEFT JOIN contract_type_master ON contract_type_master.contract_type_id = contracts.contract_type_id").
		Where("property.user_id = ? AND contract_type_master.contract_type_name = ?", userID, "amc").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// TotalActiveRentalCount returns total active rental count for a specific user
func (r *dashboardRepository) TotalActiveRentalCount(ctx context.Context, userID uint) (int64, error) {
	var count int64
	now := time.Now()
	if err := r.db.WithContext(ctx).
		Model(&models.Agreement{}).
		Joins("JOIN property ON property.property_id = agreement.property_id").
		Where("property.user_id = ? AND agreement.start_date <= ? AND agreement.end_date >= ?", userID, now, now).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// RecentActivities returns recent activities for a specific user across all domains
func (r *dashboardRepository) RecentActivities(ctx context.Context, userID uint, limit int) ([]models.ActivityItem, error) {
	var items []models.ActivityItem
	if err := r.db.WithContext(ctx).
		Table("property").
		Select("property_id as respective_id, created_at as time_created, ? as type", "property").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(limit).
		Scan(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
