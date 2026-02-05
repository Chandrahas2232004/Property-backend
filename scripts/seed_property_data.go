package main

import (
	"fmt"
	"log"
	"os"

	"property-backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("✅ Database connected")

	// Seed PropertyTypeMaster if empty
	SeedPropertyTypes(db)

	// Seed sample property data
	SeedSamplePropertyData(db)

	log.Println("✅ All seeding completed successfully")
}

// SeedPropertyTypes seeds the property_type_master table with default property types
func SeedPropertyTypes(db *gorm.DB) {
	propertyTypes := []models.PropertyTypeMaster{
		{PropertyTypeName: "residential"},
		{PropertyTypeName: "commercial"},
		{PropertyTypeName: "agricultural"},
	}

	for _, pt := range propertyTypes {
		// Check if already exists
		var existing models.PropertyTypeMaster
		if err := db.Where("property_type_name = ?", pt.PropertyTypeName).First(&existing).Error; err != nil {
			// Create if not found
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&pt).Error; err != nil {
					log.Printf("❌ Failed to seed property type %s: %v", pt.PropertyTypeName, err)
				} else {
					log.Printf("✅ Seeded property type: %s", pt.PropertyTypeName)
				}
			} else {
				log.Printf("❌ Error checking property type: %v", err)
			}
		}
	}
}

// SeedSamplePropertyData seeds sample property data with all related tables
func SeedSamplePropertyData(db *gorm.DB) {
	// Check if properties already exist
	var count int64
	db.Model(&models.Property{}).Count(&count)
	if count > 0 {
		log.Println("⏭️  Properties already seeded, skipping...")
		return
	}

	log.Println("🌱 Seeding sample property data...")

	// Get or create a sample user
	var user models.User
	if err := db.Where("email = ?", "admin@example.com").First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Get or create default role
			var role models.RolesMaster
			if err := db.Where("role = ?", "admin").First(&role).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					role = models.RolesMaster{Role: "admin"}
					db.Create(&role)
				}
			}

			user = models.User{
				FirstName:      "Admin",
				LastName:       "User",
				Email:          "admin@example.com",
				HashedPassword: "$2a$10$hashedpassword", // In real scenario, this should be properly hashed
				PhoneNumber:    "1234567890",
				RoleID:         role.RolesMasterID,
			}
			if err := db.Create(&user).Error; err != nil {
				log.Printf("❌ Failed to create sample user: %v", err)
				return
			}
		} else {
			log.Printf("❌ Error checking user: %v", err)
			return
		}
	}

	// Get property types
	var residential, commercial, agricultural models.PropertyTypeMaster
	db.Where("property_type_name = ?", "residential").First(&residential)
	db.Where("property_type_name = ?", "commercial").First(&commercial)
	db.Where("property_type_name = ?", "agricultural").First(&agricultural)

	// Get sample location data (assumes it exists)
	var taluk models.TalukMaster
	if err := db.Preload("District.State.Country").First(&taluk).Error; err != nil {
		log.Printf("⚠️  No location data found. Skipping property seeding. Please seed location data first.")
		return
	}

	// Sample properties data
	sampleProperties := []struct {
		property        models.Property
		address         models.Address
		landDetails     models.PropertyLandDetails
		taxDetails      models.PropertyTaxDetails
		ownership       models.PropertyOwnershipDetails
		buildingDetails models.PropertyBuildingDetails
		media           models.PropertyMedia
	}{
		{
			property: models.Property{
				PropertyName:   "Green Valley Estate",
				PropertyTypeID: residential.PropertyTypeID,
				UserID:         user.UserID,
				Value:          "5000000",
				Income:         "50000",
				OriginalDeed:   "yes",
			},
			address: models.Address{
				TalukID:        taluk.TalukID,
				Hobli:          "Horamavu",
				Village:        "Chelekere",
				StreetAddress:  "123 Main Street",
				Pincode:        "560067",
				LatCoordinate:  "12.9698",
				LongCoordinate: "77.7499",
			},
			landDetails: models.PropertyLandDetails{
				Rtc:        "yes",
				Ec:         "yes",
				SyNo:       "78/A",
				KhataNo:    "2023-45",
				MrNo:       "789",
				Acre:       "2.5",
				Gunte:      "100",
				Karab:      "50",
				Converted:  "yes",
				Purpose:    "residential",
				KhuskiTari: "dry",
			},
			taxDetails: models.PropertyTaxDetails{
				TaxPaid:     true,
				ReceiptNo:   "TAX2024001",
				PrevAmount:  5000.00,
				CurrAmount:  5500.00,
				ReceiptLink: "https://example.com/receipt1.pdf",
			},
			ownership: models.PropertyOwnershipDetails{
				ReceivedFrom:        "John Doe",
				AcquisitionType:     "purchase",
				RegistrationDetails: "REG2023001",
				Title:               "Absolute Owner",
				Incharge:            "Manager Name",
				PhoneNumber:         "9876543210",
			},
			buildingDetails: models.PropertyBuildingDetails{
				PlotSize:           "2500 sq.ft",
				BuiltUpArea:        "1800 sq.ft",
				YearOfConstruction: "2015",
				ApplicationNo:      "APP2015001",
			},
			media: models.PropertyMedia{
				ScannedDeedLink: "https://example.com/deed1.pdf",
				PhotoLink:       "https://example.com/photo1.jpg",
				Remarks:         "Property in excellent condition",
			},
		},
		{
			property: models.Property{
				PropertyName:   "Sunshine Commercial Complex",
				PropertyTypeID: commercial.PropertyTypeID,
				UserID:         user.UserID,
				Value:          "8000000",
				Income:         "100000",
				OriginalDeed:   "yes",
			},
			address: models.Address{
				TalukID:        taluk.TalukID,
				Hobli:          "Marathahalli",
				Village:        "Kadubeesanahalli",
				StreetAddress:  "456 Business Road",
				Pincode:        "560037",
				LatCoordinate:  "12.9352",
				LongCoordinate: "77.6975",
			},
			landDetails: models.PropertyLandDetails{
				Rtc:        "yes",
				Ec:         "yes",
				SyNo:       "92/B",
				KhataNo:    "2022-88",
				MrNo:       "456",
				Acre:       "1.0",
				Gunte:      "40",
				Karab:      "40",
				Converted:  "yes",
				Purpose:    "commercial",
				KhuskiTari: "irrigated",
			},
			taxDetails: models.PropertyTaxDetails{
				TaxPaid:     true,
				ReceiptNo:   "TAX2024002",
				PrevAmount:  8000.00,
				CurrAmount:  8500.00,
				ReceiptLink: "https://example.com/receipt2.pdf",
			},
			ownership: models.PropertyOwnershipDetails{
				ReceivedFrom:        "ABC Corporation",
				AcquisitionType:     "lease",
				RegistrationDetails: "REG2022050",
				Title:               "Leaseholder",
				Incharge:            "Property Manager",
				PhoneNumber:         "9876543211",
			},
			buildingDetails: models.PropertyBuildingDetails{
				PlotSize:           "4000 sq.ft",
				BuiltUpArea:        "3500 sq.ft",
				YearOfConstruction: "2018",
				ApplicationNo:      "APP2018025",
			},
			media: models.PropertyMedia{
				ScannedDeedLink: "https://example.com/deed2.pdf",
				PhotoLink:       "https://example.com/photo2.jpg",
				Remarks:         "Prime location for business",
			},
		},
		{
			property: models.Property{
				PropertyName:   "Farmland Paradise",
				PropertyTypeID: agricultural.PropertyTypeID,
				UserID:         user.UserID,
				Value:          "3000000",
				Income:         "25000",
				OriginalDeed:   "yes",
			},
			address: models.Address{
				TalukID:        taluk.TalukID,
				Hobli:          "Jala",
				Village:        "Kengeri",
				StreetAddress:  "Farm Road 789",
				Pincode:        "560060",
				LatCoordinate:  "12.9081",
				LongCoordinate: "77.4852",
			},
			landDetails: models.PropertyLandDetails{
				Rtc:        "yes",
				Ec:         "yes",
				SyNo:       "34/C",
				KhataNo:    "2020-12",
				MrNo:       "234",
				Acre:       "5.0",
				Gunte:      "200",
				Karab:      "180",
				Converted:  "no",
				Purpose:    "agriculture",
				KhuskiTari: "irrigated",
			},
			taxDetails: models.PropertyTaxDetails{
				TaxPaid:     true,
				ReceiptNo:   "TAX2024003",
				PrevAmount:  3000.00,
				CurrAmount:  3200.00,
				ReceiptLink: "https://example.com/receipt3.pdf",
			},
			ownership: models.PropertyOwnershipDetails{
				ReceivedFrom:        "Ancestral Property",
				AcquisitionType:     "inheritance",
				RegistrationDetails: "REG2020123",
				Title:               "Owner",
				Incharge:            "Farm Manager",
				PhoneNumber:         "9876543212",
			},
			buildingDetails: models.PropertyBuildingDetails{
				PlotSize:           "217800 sq.ft",
				BuiltUpArea:        "500 sq.ft",
				YearOfConstruction: "2010",
				ApplicationNo:      "APP2010010",
			},
			media: models.PropertyMedia{
				ScannedDeedLink: "https://example.com/deed3.pdf",
				PhotoLink:       "https://example.com/photo3.jpg",
				Remarks:         "Fertile land with water source",
			},
		},
	}

	// Insert sample data
	for i, sample := range sampleProperties {
		tx := db.Begin()

		// Create address
		if err := tx.Create(&sample.address).Error; err != nil {
			tx.Rollback()
			log.Printf("❌ Failed to create address for property %d: %v", i+1, err)
			continue
		}

		// Create property
		sample.property.AddressID = sample.address.AddressID
		if err := tx.Create(&sample.property).Error; err != nil {
			tx.Rollback()
			log.Printf("❌ Failed to create property %d: %v", i+1, err)
			continue
		}

		// Create land details
		sample.landDetails.PropertyID = sample.property.PropertyID
		if err := tx.Create(&sample.landDetails).Error; err != nil {
			tx.Rollback()
			log.Printf("❌ Failed to create land details for property %d: %v", i+1, err)
			continue
		}

		// Create tax details
		sample.taxDetails.PropertyID = sample.property.PropertyID
		if err := tx.Create(&sample.taxDetails).Error; err != nil {
			tx.Rollback()
			log.Printf("❌ Failed to create tax details for property %d: %v", i+1, err)
			continue
		}

		// Create ownership details
		sample.ownership.PropertyID = sample.property.PropertyID
		if err := tx.Create(&sample.ownership).Error; err != nil {
			tx.Rollback()
			log.Printf("❌ Failed to create ownership details for property %d: %v", i+1, err)
			continue
		}

		// Create building details
		sample.buildingDetails.PropertyID = sample.property.PropertyID
		if err := tx.Create(&sample.buildingDetails).Error; err != nil {
			tx.Rollback()
			log.Printf("❌ Failed to create building details for property %d: %v", i+1, err)
			continue
		}

		// Create media
		sample.media.PropertyID = sample.property.PropertyID
		if err := tx.Create(&sample.media).Error; err != nil {
			tx.Rollback()
			log.Printf("❌ Failed to create media for property %d: %v", i+1, err)
			continue
		}

		// Commit transaction
		if err := tx.Commit().Error; err != nil {
			log.Printf("❌ Failed to commit property %d: %v", i+1, err)
			continue
		}

		log.Printf("✅ Seeded property: %s", sample.property.PropertyName)
	}

	log.Println("✅ Sample property data seeding completed")
}
