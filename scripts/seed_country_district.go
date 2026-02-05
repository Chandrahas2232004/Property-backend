package main

import (
    "fmt"
    "log"
    "os"
    "strings"

    "property-backend/models"

    "github.com/joho/godotenv"
    "github.com/xuri/excelize/v2"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func mustLoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println(".env not found or failed to load — proceeding with environment variables")
    }
}

func openDB() *gorm.DB {
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
        log.Fatalf("failed to open db: %v", err)
    }
    return db
}

func main() {
    mustLoadEnv()
    db := openDB()

    filePath := "locationData/villages-directory.xlsx"
    f, err := excelize.OpenFile(filePath)
    if err != nil {
        log.Fatalf("failed to open excel file %s: %v", filePath, err)
    }

    rows, err := f.GetRows("village-directory")
    if err != nil {
        log.Fatalf("failed to read sheet 'village-directory': %v", err)
    }

    // ensure country
    var country models.CountryMaster
    if err := db.FirstOrCreate(&country, models.CountryMaster{CountryName: "India"}).Error; err != nil {
        log.Fatalf("failed to ensure country: %v", err)
    }

    created := 0
    skipped := 0

    for i, row := range rows {
        if i == 0 || len(row) < 5 {
            continue
        }

        stateName := strings.TrimSpace(row[1])
        districtName := strings.TrimSpace(row[4])

        if stateName == "" || districtName == "" {
            skipped++
            continue
        }

        // find state (states are assumed already seeded)
        var state models.StateMaster
        if err := db.Where("state_name = ? AND country_id = ?", stateName, country.CountryID).First(&state).Error; err != nil {
            // state not found — skip
            log.Printf("state not found for row %d: %s — skipping district %s", i+1, stateName, districtName)
            skipped++
            continue
        }

        // create district if not exists
        var district models.DistrictMaster
        if err := db.FirstOrCreate(&district, models.DistrictMaster{DistrictName: districtName, StateID: state.StateID}).Error; err != nil {
            log.Printf("failed to create district '%s' for state '%s': %v", districtName, stateName, err)
            skipped++
            continue
        }

        created++
    }

    log.Printf("✅ Country and district seeding completed — created: %d, skipped: %d\n", created, skipped)
}
