package main

import (
    "log"

    "property-backend/config"
)

func main() {
    // ConnectDatabase will auto-migrate and (already) call SeedLocationFromExcel,
    // but call it explicitly to ensure seeding runs in this standalone script.
    config.ConnectDatabase()

    // Ensure seed is executed with configured DB
    config.SeedLocationFromExcel(config.DB)

    log.Println("✅ Location seeding completed")
}
