package models

import "time"

/* =========================
   agreement
========================= */

type Agreement struct {
	AgreementID uint      `gorm:"column:agreement_id;primaryKey;autoIncrement" json:"agreement_id"`
	PropertyID  uint      `gorm:"column:property_id;not null" json:"property_id"`
	TenantName  string    `gorm:"column:tenant_name;type:varchar(150);not null" json:"tenant_name"`
	ContactNo   string    `gorm:"column:contact_no;type:varchar(15)" json:"contact_no"`
	Rent        float64   `gorm:"column:rent;not null" json:"rent"`
	Deposit     float64   `gorm:"column:deposit" json:"deposit"`
	StartDate   time.Time `gorm:"column:start_date;not null" json:"start_date"`
	EndDate     time.Time `gorm:"column:end_date;not null" json:"end_date"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`

	/* relation */
	Property Property `gorm:"foreignKey:PropertyID;references:PropertyID" json:"property"`
}

func (Agreement) TableName() string {
	return "agreement"
}
