package models

import "time"

/* =========================
   asset_type_master
========================= */

type AssetTypeMaster struct {
	AssetTypeID   uint   `gorm:"column:asset_type_id;primaryKey;autoIncrement"`
	AssetTypeName string `gorm:"column:asset_type_name;type:varchar(100);not null;unique"`
}

func (AssetTypeMaster) TableName() string {
	return "asset_type_master"
}

/* =========================
   assets
========================= */

type Assets struct {
	AssetID   uint   `gorm:"column:asset_id;primaryKey;autoIncrement"`
	AssetName string `gorm:"column:asset_name;type:varchar(150);not null"`

	/* foreign key → asset_type_master */
	AssetTypeID uint `gorm:"column:asset_type_id;not null"`

	AssetLocation    string `gorm:"column:location;type:varchar(150)"`
	AssetCost        string `gorm:"column:cost;type:varchar(50)"`
	AssetAMCProvider string `gorm:"column:amc_provider;type:varchar(150)"`
	AssetStartDate   string `gorm:"column:start_date;type:varchar(20)"`
	AssetEndDate     string `gorm:"column:end_date;type:varchar(20)"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`

	/* foreign key → property */
	PropertyID uint `gorm:"column:property_id;not null"`

	/* relations */
	Property  Property        `gorm:"foreignKey:PropertyID;references:PropertyID"`
	AssetType AssetTypeMaster `gorm:"foreignKey:AssetTypeID;references:AssetTypeID"`
}

func (Assets) TableName() string {
	return "assets"
}
