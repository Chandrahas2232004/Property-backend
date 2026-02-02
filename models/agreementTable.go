package models

import "time"

/* =========================
   agreement
========================= */

type Agreement struct {
	AgreementID uint      `gorm:"column:agreement_id;primaryKey;autoIncrement"`
	PropertyID  uint      `gorm:"column:property_id;not null"`
	TenantName  string    `gorm:"column:tenant_name;type:varchar(150);not null"`
	ContactNo   string    `gorm:"column:contact_no;type:varchar(15)"`
	Rent        float64   `gorm:"column:rent;not null"`
	Deposit     float64   `gorm:"column:deposit"`
	StartDate   time.Time `gorm:"column:start_date;not null"`
	EndDate     time.Time `gorm:"column:end_date;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`

	/* relation */
	Property Property `gorm:"foreignKey:PropertyID;references:PropertyID"`
}

func (Agreement) TableName() string {
	return "agreement"
}
