package repositories

import (
	"context"

	"property-backend/models"

	"gorm.io/gorm"
)

// FormData represents ID and Name pair for form options
type FormData struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// FormRepository defines form-related data access methods
type FormRepository interface {
	GetCountry(ctx context.Context) ([]FormData, error)
	GetState(ctx context.Context, countryID string) ([]FormData, error)
	GetDistrict(ctx context.Context, stateID string) ([]FormData, error)
	GetTaluk(ctx context.Context, districtID string) ([]FormData, error)
}

type formRepository struct {
	db *gorm.DB
}

// NewFormRepository constructs a FormRepository
func NewFormRepository(db *gorm.DB) FormRepository {
	return &formRepository{db: db}
}

// GetCountry retrieves all countries from the database
func (r *formRepository) GetCountry(ctx context.Context) ([]FormData, error) {
	var countries []models.CountryMaster
	err := r.db.WithContext(ctx).
		Model(&models.CountryMaster{}).
		Select("country_id, country_name").
		Find(&countries).Error
	if err != nil {
		return nil, err
	}

	var result []FormData
	for _, country := range countries {
		result = append(result, FormData{
			ID:   country.CountryID,
			Name: country.CountryName,
		})
	}
	return result, nil
}

// GetState retrieves all states for a given country ID
func (r *formRepository) GetState(ctx context.Context, countryID string) ([]FormData, error) {
	var states []models.StateMaster
	err := r.db.WithContext(ctx).
		Model(&models.StateMaster{}).
		Select("state_id, state_name").
		Where("country_id = ?", countryID).
		Find(&states).Error
	if err != nil {
		return nil, err
	}

	var result []FormData
	for _, state := range states {
		result = append(result, FormData{
			ID:   state.StateID,
			Name: state.StateName,
		})
	}
	return result, nil
}

// GetDistrict retrieves all districts for a given state ID
func (r *formRepository) GetDistrict(ctx context.Context, stateID string) ([]FormData, error) {
	var districts []models.DistrictMaster
	err := r.db.WithContext(ctx).
		Model(&models.DistrictMaster{}).
		Select("district_id, district_name").
		Where("state_id = ?", stateID).
		Find(&districts).Error
	if err != nil {
		return nil, err
	}

	var result []FormData
	for _, district := range districts {
		result = append(result, FormData{
			ID:   district.DistrictID,
			Name: district.DistrictName,
		})
	}
	return result, nil
}

// GetTaluk retrieves all taluks for a given district ID
func (r *formRepository) GetTaluk(ctx context.Context, districtID string) ([]FormData, error) {
	var taluks []models.TalukMaster
	err := r.db.WithContext(ctx).
		Model(&models.TalukMaster{}).
		Select("taluk_id, taluk_name").
		Where("district_id = ?", districtID).
		Find(&taluks).Error
	if err != nil {
		return nil, err
	}

	var result []FormData
	for _, taluk := range taluks {
		result = append(result, FormData{
			ID:   taluk.TalukID,
			Name: taluk.TalukName,
		})
	}
	return result, nil
}
