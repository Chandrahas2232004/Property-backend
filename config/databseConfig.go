package config

import (
	"fmt"
	"log"
	"os"

	"property-backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
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

	// ✅ OPEN DB ONLY ONCE
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("✅ Database connected")

	// ✅ MIGRATION ORDER MATTERS
	err = DB.AutoMigrate(

		// masters
		&models.RolesMaster{},
		&models.CountryMaster{},
		&models.StateMaster{},
		&models.DistrictMaster{},
		&models.TalukMaster{},
		&models.Address{},
		&models.PropertyTypeMaster{},
		&models.AssetTypeMaster{},
		&models.ContractTypeMaster{},

		// core tables
		&models.User{},
		&models.Property{},
		&models.Asset{},
		&models.Contract{},
		&models.Agreement{},

		// property related child tables
		&models.UserRoles{},
		&models.PropertyLandDetails{},
		&models.PropertyBuildingDetails{},
		&models.PropertyTaxDetails{},
		&models.PropertyOwnershipDetails{},
		&models.PropertyMedia{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("✅ Tables migrated successfully")
}
