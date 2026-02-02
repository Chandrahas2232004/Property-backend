package models

import "time"

/* =========================
   contract_type_master
========================= */

type ContractTypeMaster struct {
	ContractTypeID   uint   `gorm:"column:contract_type_id;primaryKey;autoIncrement"`
	ContractTypeName string `gorm:"column:contract_type_name;type:varchar(100);not null;unique"`
}

func (ContractTypeMaster) TableName() string {
	return "contract_type_master"
}

/* =========================
   contracts
========================= */

type Contracts struct {
	ContractID uint `gorm:"column:contract_id;primaryKey;autoIncrement"`

	Name string `gorm:"column:name;type:varchar(150);not null"`

	/* foreign key → contract_type_master */
	ContractTypeID uint `gorm:"column:contract_type_id;not null"`

	RelatedTo string  `gorm:"column:related_to;type:varchar(50)"`
	Cost      float64 `gorm:"column:cost"`
	Provider  string  `gorm:"column:provider;type:varchar(150)"`

	StartDate time.Time `gorm:"column:start_date"`
	EndDate   time.Time `gorm:"column:end_date"`

	Terms string `gorm:"column:terms;type:text"`

	/* foreign key → assets */
	AssetID uint `gorm:"column:asset_id;not null"`

	/* relations */
	Asset        Assets             `gorm:"foreignKey:AssetID;references:AssetID"`
	ContractType ContractTypeMaster `gorm:"foreignKey:ContractTypeID;references:ContractTypeID"`
}

func (Contracts) TableName() string {
	return "contracts"
}
