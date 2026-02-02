package config

import (
	"log"
	"strings"

	"property-backend/models"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func SeedLocationFromExcel(db *gorm.DB) {

	filePath := "locationData/villages-directory.xlsx"

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal("Failed to open excel file:", err)
	}

	rows, err := f.GetRows("village-directory")
	if err != nil {
		log.Fatal("Failed to read sheet:", err)
	}

	// ---- Ensure Country ----
	var country models.CountryMaster
	db.FirstOrCreate(&country, models.CountryMaster{
		CountryName: "India",
	})

	// ---- Loop rows (skip header) ----
	for i, row := range rows {
		if i == 0 || len(row) < 8 {
			continue
		}

		stateName := strings.TrimSpace(row[1])   // State Name
		districtName := strings.TrimSpace(row[4]) // District Name
		talukName := strings.TrimSpace(row[7])    // Subdistrict Name

		if stateName == "" || districtName == "" || talukName == "" {
			continue
		}

		// ---- State ----
		var state models.StateMaster
		db.FirstOrCreate(&state, models.StateMaster{
			StateName: stateName,
			CountryID: country.CountryID,
		})

		// ---- District ----
		var district models.DistrictMaster
		db.FirstOrCreate(&district, models.DistrictMaster{
			DistrictName: districtName,
			StateID:      state.StateID,
		})

		// ---- Taluk ----
		var taluk models.TalukMaster
		db.FirstOrCreate(&taluk, models.TalukMaster{
			TalukName:  talukName,
			DistrictID: district.DistrictID,
		})
	}

	log.Println("✅ Excel census data imported successfully")
}
