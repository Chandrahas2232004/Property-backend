package repositories

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
	"property-backend/models"
)

// PropertyRepository defines property-related data access methods
type PropertyRepository interface {
	Total(ctx context.Context) (int64, error)
	ActiveRentalCount(ctx context.Context) (int64, error)
	Create(ctx context.Context, req interface{}) (int64, error)
	ListByType(ctx context.Context, propertyType string) ([]models.Property, error)
}

type propertyRepository struct {
	db *gorm.DB
}

// NewPropertyRepository constructs a PropertyRepository
func NewPropertyRepository(db *gorm.DB) PropertyRepository {
	return &propertyRepository{db: db}
}

func (r *propertyRepository) Total(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Property{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *propertyRepository) ActiveRentalCount(ctx context.Context) (int64, error) {
	var count int64
	now := time.Now()
	if err := r.db.WithContext(ctx).
		Model(&models.Agreement{}).
		Where("start_date <= ? AND end_date >= ?", now, now).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *propertyRepository) Create(ctx context.Context, req interface{}) (int64, error) {
	// Type assert to get the request object
	requestData := req.(map[string]interface{})

	// Start transaction
	tx := r.db.WithContext(ctx).Begin()

	// 1. Get or create PropertyType
	propertyTypeID := uint(0)
	if ptID, ok := requestData["property_type_id"].(float64); ok && ptID > 0 {
		propertyTypeID = uint(ptID)
	} else if ptName, ok := requestData["property_type_name"].(string); ok && ptName != "" {
		var propType models.PropertyTypeMaster
		if err := tx.Where("property_type_name = ?", ptName).First(&propType).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				propType.PropertyTypeName = ptName
				if err := tx.Create(&propType).Error; err != nil {
					tx.Rollback()
					return 0, fmt.Errorf("failed to create property type: %w", err)
				}
			} else {
				tx.Rollback()
				return 0, err
			}
		}
		propertyTypeID = propType.PropertyTypeID
	}

	// 2. Create Address hierarchy (Country -> State -> District -> Taluk -> Address)
	var countryID, stateID, districtID, talukID, addressID uint

	// Get or create Country
	if countryName, ok := requestData["country_name"].(string); ok && countryName != "" {
		var country models.CountryMaster
		if err := tx.Where("country_name = ?", countryName).First(&country).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				country.CountryName = countryName
				if err := tx.Create(&country).Error; err != nil {
					tx.Rollback()
					return 0, fmt.Errorf("failed to create country: %w", err)
				}
			} else {
				tx.Rollback()
				return 0, err
			}
		}
		countryID = country.CountryID
	}

	// Get or create State
	if stateName, ok := requestData["state_name"].(string); ok && stateName != "" && countryID > 0 {
		var state models.StateMaster
		if err := tx.Where("state_name = ? AND country_id = ?", stateName, countryID).First(&state).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				state.StateName = stateName
				state.CountryID = countryID
				if err := tx.Create(&state).Error; err != nil {
					tx.Rollback()
					return 0, fmt.Errorf("failed to create state: %w", err)
				}
			} else {
				tx.Rollback()
				return 0, err
			}
		}
		stateID = state.StateID
	}

	// Get or create District
	if districtName, ok := requestData["district_name"].(string); ok && districtName != "" && stateID > 0 {
		var district models.DistrictMaster
		if err := tx.Where("district_name = ? AND state_id = ?", districtName, stateID).First(&district).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				district.DistrictName = districtName
				district.StateID = stateID
				if err := tx.Create(&district).Error; err != nil {
					tx.Rollback()
					return 0, fmt.Errorf("failed to create district: %w", err)
				}
			} else {
				tx.Rollback()
				return 0, err
			}
		}
		districtID = district.DistrictID
	}

	// Get or create Taluk
	if talukName, ok := requestData["taluk_name"].(string); ok && talukName != "" && districtID > 0 {
		var taluk models.TalukMaster
		if err := tx.Where("taluk_name = ? AND district_id = ?", talukName, districtID).First(&taluk).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				taluk.TalukName = talukName
				taluk.DistrictID = districtID
				if err := tx.Create(&taluk).Error; err != nil {
					tx.Rollback()
					return 0, fmt.Errorf("failed to create taluk: %w", err)
				}
			} else {
				tx.Rollback()
				return 0, err
			}
		}
		talukID = taluk.TalukID
	}

	// Create Address
	if talukID > 0 {
		address := models.Address{
			TalukID:        talukID,
			Hobli:          getStringValue(requestData, "hobli"),
			Village:        getStringValue(requestData, "village"),
			StreetAddress:  getStringValue(requestData, "street_address"),
			Pincode:        getStringValue(requestData, "pincode"),
			LatCoordinate:  getStringValue(requestData, "lat_coordinate"),
			LongCoordinate: getStringValue(requestData, "long_coordinate"),
		}
		if err := tx.Create(&address).Error; err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("failed to create address: %w", err)
		}
		addressID = address.AddressID
	}

	// 3. Create Property
	property := models.Property{
		PropertyName:   getStringValue(requestData, "property_name"),
		PropertyTypeID: propertyTypeID,
		AddressID:      addressID,
		UserID:         uint(getFloatValue(requestData, "user_id")),
		Value:          getStringValue(requestData, "value"),
		Income:         getStringValue(requestData, "income"),
		OriginalDeed:   getStringValue(requestData, "original_deed"),
	}

	if err := tx.Create(&property).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create property: %w", err)
	}

	// 4. Create PropertyLandDetails
	landDetails := models.PropertyLandDetails{
		PropertyID: property.PropertyID,
		Rtc:        getStringValue(requestData, "rtc"),
		Ec:         getStringValue(requestData, "ec"),
		SyNo:       getStringValue(requestData, "syno"),
		KhataNo:    getStringValue(requestData, "khatano"),
		MrNo:       getStringValue(requestData, "mrno"),
		Acre:       getStringValue(requestData, "acre"),
		Gunte:      getStringValue(requestData, "gunte"),
		Karab:      getStringValue(requestData, "karab"),
		Converted:  getStringValue(requestData, "converted"),
		Purpose:    getStringValue(requestData, "purpose"),
		KhuskiTari: getStringValue(requestData, "khuskitari"),
	}
	if err := tx.Create(&landDetails).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create land details: %w", err)
	}

	// 5. Create PropertyTaxDetails
	taxDetails := models.PropertyTaxDetails{
		PropertyID:  property.PropertyID,
		ReceiptNo:   getStringValue(requestData, "receipt_no"),
		PrevAmount:  getFloatValue(requestData, "prev_amount"),
		CurrAmount:  getFloatValue(requestData, "curr_amount"),
		ReceiptLink: getStringValue(requestData, "receipt_link"),
	}
	if err := tx.Create(&taxDetails).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create tax details: %w", err)
	}

	// 6. Create PropertyOwnershipDetails
	ownershipDetails := models.PropertyOwnershipDetails{
		PropertyID:          property.PropertyID,
		ReceivedFrom:        getStringValue(requestData, "received_from"),
		AcquisitionType:     getStringValue(requestData, "acquisition_type"),
		RegistrationDetails: getStringValue(requestData, "registration_details"),
		Title:               getStringValue(requestData, "title"),
		Incharge:            getStringValue(requestData, "incharge"),
		PhoneNumber:         getStringValue(requestData, "phone_number"),
	}
	if err := tx.Create(&ownershipDetails).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create ownership details: %w", err)
	}

	// 7. Create PropertyBuildingDetails
	buildingDetails := models.PropertyBuildingDetails{
		PropertyID:         property.PropertyID,
		PlotSize:           getStringValue(requestData, "plot_size"),
		BuiltUpArea:        getStringValue(requestData, "built_up_area"),
		YearOfConstruction: getStringValue(requestData, "year_of_construction"),
		ApplicationNo:      getStringValue(requestData, "application_no"),
	}
	if err := tx.Create(&buildingDetails).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create building details: %w", err)
	}

	// 8. Create PropertyMedia
	media := models.PropertyMedia{
		PropertyID:      property.PropertyID,
		ScannedDeedLink: getStringValue(requestData, "scanned_deed_link"),
		PhotoLink:       getStringValue(requestData, "photo_link"),
		Remarks:         getStringValue(requestData, "remarks"),
	}
	if err := tx.Create(&media).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create media: %w", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return int64(property.PropertyID), nil
}

func (r *propertyRepository) ListByType(ctx context.Context, propertyType string) ([]models.Property, error) {
	var props []models.Property
	if err := r.db.WithContext(ctx).
		Preload("Address").
		Preload("PropertyType").
		Joins("JOIN property_type_master ON property_type_master.property_type_id = property.property_type_id").
		Where("property_type_master.property_type_name = ?", propertyType).
		Find(&props).Error; err != nil {
		return nil, err
	}
	return props, nil
}

// Helper functions
func getStringValue(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getFloatValue(data map[string]interface{}, key string) float64 {
	if val, ok := data[key]; ok {
		if f, ok := val.(float64); ok {
			return f
		}
		if str, ok := val.(string); ok {
			if f, err := strconv.ParseFloat(str, 64); err == nil {
				return f
			}
		}
	}
	return 0
}
