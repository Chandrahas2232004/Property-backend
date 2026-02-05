package models

import "time"

/* =========================
   asset_type_master
========================= */

type AssetTypeMaster struct {
	AssetTypeID   uint   `gorm:"column:asset_type_id;primaryKey;autoIncrement" json:"asset_type_id"`
	AssetTypeName string `gorm:"column:asset_type_name;type:varchar(100);not null;unique" json:"asset_type_name"`
}

func (AssetTypeMaster) TableName() string {
	return "asset_type_master"
}

/* =========================
   assets
========================= */

type Asset struct {
	AssetID   uint   `gorm:"column:asset_id;primaryKey;autoIncrement" json:"asset_id"`
	AssetName string `gorm:"column:asset_name;type:varchar(150);not null" json:"asset_name"`

	/* foreign key → asset_type_master */
	AssetTypeID uint `gorm:"column:asset_type_id;not null" json:"asset_type_id"`

	AssetLocation    string `gorm:"column:location;type:varchar(150)" json:"asset_location"`
	AssetCost        string `gorm:"column:cost;type:varchar(50)" json:"asset_cost"`
	AssetAMCProvider string `gorm:"column:amc_provider;type:varchar(150)" json:"asset_amc_provider"`
	AssetStartDate   string `gorm:"column:start_date;type:varchar(20)" json:"asset_start_date"`
	AssetEndDate     string `gorm:"column:end_date;type:varchar(20)" json:"asset_end_date"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`

	/* foreign key → property */
	PropertyID uint `gorm:"column:property_id;not null" json:"property_id"`

	/* relations */
	Property  Property        `gorm:"foreignKey:PropertyID;references:PropertyID" json:"property"`
	AssetType AssetTypeMaster `gorm:"foreignKey:AssetTypeID;references:AssetTypeID" json:"asset_type"`
}

func (Asset) TableName() string {
	return "assets"
}
