package repositories

import (
	"context"

	"gorm.io/gorm"
	"property-backend/models"
)

// FormRepository defines form-related data access methods
type FormRepository interface {
	GetCountry(ctx context.Context) ([]string, error)
	GetState(ctx context.Context, countryID string) ([]string, error)
	GetDistrict(ctx context.Context, stateID string) ([]string, error)
	GetTaluk(ctx context.Context, districtID string) ([]string, error)
}

type formRepository struct {
	db *gorm.DB
}

// NewFormRepository constructs a FormRepository
func NewFormRepository(db *gorm.DB) FormRepository {
	return &formRepository{db: db}
}

// GetCountry retrieves all countries from the database
func (r *formRepository) GetCountry(ctx context.Context) ([]string, error) {
	var countryNames []string
	err := r.db.WithContext(ctx).
		Model(&models.CountryMaster{}).
		Pluck("country_name", &countryNames).Error
	if err != nil {
		return nil, err
	}
	return countryNames, nil
}

// GetState retrieves all states for a given country ID
func (r *formRepository) GetState(ctx context.Context, countryID string) ([]string, error) {
	var stateNames []string
	err := r.db.WithContext(ctx).
		Model(&models.StateMaster{}).
		Where("country_id = ?", countryID).
		Pluck("state_name", &stateNames).Error
	if err != nil {
		return nil, err
	}
	return stateNames, nil
}

// GetDistrict retrieves all districts for a given state ID
func (r *formRepository) GetDistrict(ctx context.Context, stateID string) ([]string, error) {
	var districtNames []string
	err := r.db.WithContext(ctx).
		Model(&models.DistrictMaster{}).
		Where("state_id = ?", stateID).
		Pluck("district_name", &districtNames).Error
	if err != nil {
		return nil, err
	}
	return districtNames, nil
}

// GetTaluk retrieves all taluks for a given district ID
func (r *formRepository) GetTaluk(ctx context.Context, districtID string) ([]string, error) {
	var talukNames []string
	err := r.db.WithContext(ctx).
		Model(&models.TalukMaster{}).
		Where("district_id = ?", districtID).
		Pluck("taluk_name", &talukNames).Error
	if err != nil {
		return nil, err
	}
	return talukNames, nil
}