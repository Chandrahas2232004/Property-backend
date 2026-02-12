# Seed Scripts

These scripts are designed to run independently to seed your database with initial data. Each script has its own `main()` function and should be run separately.

## Running the Scripts

Run each script from the project root directory:

### 1. Seed Location Data
Seeds country, state, district, and taluk from Excel file:
```bash
go run ./scripts/seed_location/main.go
```
**Prerequisites:** Location CSV/Excel file must exist at `locationData/village-directory.csv`

### 2. Seed Country & District
Populates country and district master tables from Excel:
```bash
go run ./scripts/seed_country_district/main.go
```
**Requires:** `locationData/villages-directory.xlsx` file

### 3. Seed Property Data
Creates sample property records with all related details:
```bash
go run ./scripts/seed_property_data/main.go
```
**Requires:** Location data (run seed_location first) and Database connection configured in `.env`

## Recommended Execution Order

1. `seed_location/main.go` - Initialize location hierarchy
2. `seed_country_district/main.go` - Populate country/district if needed
3. `seed_property_data/main.go` - Create sample property records

## Environment Setup

Ensure your `.env` file contains:
```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=property_db
DB_PORT=5432
DB_SSLMODE=disable
```

## Notes

- Each script connects to the database independently
- Scripts check for existing data to avoid duplicates
- Old script files in `scripts/` root are deprecated - use the subdirectories instead
