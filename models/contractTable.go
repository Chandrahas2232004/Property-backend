package models

import "time"

/* =========================
   contract_type_master
========================= */

type ContractTypeMaster struct {
	ContractTypeID   uint   `gorm:"column:contract_type_id;primaryKey;autoIncrement" json:"contract_type_id"`
	ContractTypeName string `gorm:"column:contract_type_name;type:varchar(100);not null;unique" json:"contract_type_name"`
}

func (ContractTypeMaster) TableName() string {
	return "contract_type_master"
}

/* =========================
   contracts
========================= */

type Contract struct {
	ContractID uint `gorm:"column:contract_id;primaryKey;autoIncrement" json:"contract_id"`

	Name string `gorm:"column:name;type:varchar(150);not null" json:"name"`

	/* foreign key → contract_type_master */
	ContractTypeID uint `gorm:"column:contract_type_id;not null" json:"contract_type_id"`

	RelatedTo string  `gorm:"column:related_to;type:varchar(50)" json:"related_to"`
	Cost      float64 `gorm:"column:cost" json:"cost"`
	Provider  string  `gorm:"column:provider;type:varchar(150)" json:"provider"`

	StartDate time.Time `gorm:"column:start_date" json:"start_date"`
	EndDate   time.Time `gorm:"column:end_date" json:"end_date"`

	Terms string `gorm:"column:terms;type:text" json:"terms"`

	/* foreign key → assets */
	AssetID uint `gorm:"column:asset_id;not null" json:"asset_id"`

	/* relations */
	Asset        Asset              `gorm:"foreignKey:AssetID;references:AssetID" json:"asset"`
	ContractType ContractTypeMaster `gorm:"foreignKey:ContractTypeID;references:ContractTypeID" json:"contract_type"`
}

func (Contract) TableName() string {
	return "contracts"
}
