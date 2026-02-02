package models

import "time"

/* =========================
   country_master
========================= */

type CountryMaster struct {
	CountryID   uint   `gorm:"column:country_id;primaryKey;autoIncrement"`
	CountryName string `gorm:"column:country_name;type:varchar(100);not null;unique"`
}

/* =========================
   state_master
========================= */

type StateMaster struct {
	StateID   uint   `gorm:"column:state_id;primaryKey;autoIncrement"`
	StateName string `gorm:"column:state_name;type:varchar(100);not null"`
	CountryID uint   `gorm:"column:country_id;not null"`

	Country CountryMaster `gorm:"foreignKey:CountryID;references:CountryID"`
}

/* =========================
   district_master
========================= */

type DistrictMaster struct {
	DistrictID   uint   `gorm:"column:district_id;primaryKey;autoIncrement"`
	DistrictName string `gorm:"column:district_name;type:varchar(100);not null"`
	StateID      uint   `gorm:"column:state_id;not null"`

	State StateMaster `gorm:"foreignKey:StateID;references:StateID"`
}

/* =========================
   taluk_master
========================= */

type TalukMaster struct {
	TalukID    uint   `gorm:"column:taluk_id;primaryKey;autoIncrement"`
	TalukName  string `gorm:"column:taluk_name;type:varchar(100);not null"`
	DistrictID uint   `gorm:"column:district_id;not null"`

	District DistrictMaster `gorm:"foreignKey:DistrictID;references:DistrictID"`
}

/* =========================
   address
========================= */

type Address struct {
	AddressID      uint   `gorm:"column:address_id;primaryKey;autoIncrement"`
	TalukID        uint   `gorm:"column:taluk_id;not null"`
	Hobli          string `gorm:"column:hobli;type:varchar(100)"`
	Village        string `gorm:"column:village;type:varchar(100)"`
	StreetAddress  string `gorm:"column:street_address;type:text"`
	Pincode        string `gorm:"column:pincode;type:varchar(10)"`
	LatCoordinate  string `gorm:"column:lat_coordinate;type:varchar(50)"`
	LongCoordinate string `gorm:"column:long_coordinate;type:varchar(50)"`

	Taluk TalukMaster `gorm:"foreignKey:TalukID;references:TalukID"`
}

/* =========================
   property
========================= */

type Property struct {
	PropertyID     uint      `gorm:"column:property_id;primaryKey;autoIncrement"`
	PropertyName   string    `gorm:"column:property_name;type:varchar(150);not null"`
	PropertyTypeID uint      `gorm:"column:property_type_id;not null"`

	Value        string    `gorm:"column:value;type:varchar(50)"`
	Description  string    `gorm:"column:description;type:text"`
	Purpose      string    `gorm:"column:purpose;type:varchar(100)"`
	Income       string    `gorm:"column:income;type:varchar(50)"`
	OriginalDeed string    `gorm:"column:original_deed;type:varchar(10)"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`

	AddressID uint `gorm:"column:address_id;not null"`
	UserID    uint `gorm:"column:user_id;not null"`

	/* relations (FKs live ONLY here — correct direction) */
	Address      Address            `gorm:"foreignKey:AddressID;references:AddressID;constraint:OnUpdate:CASCADE"`
	User         User               `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE"`
	PropertyType PropertyTypeMaster `gorm:"foreignKey:PropertyTypeID;references:PropertyTypeID;constraint:OnUpdate:CASCADE"`
}

func (Property) TableName() string {
	return "property"
}


/* =========================
   property_type_master
========================= */

type PropertyTypeMaster struct {
	PropertyTypeID   uint   `gorm:"column:property_type_id;primaryKey;autoIncrement"`
	PropertyTypeName string `gorm:"column:property_type_name;type:varchar(100);not null;unique"`
}

/* =========================
   property_land_details
========================= */

type PropertyLandDetails struct {
	LandDetailsID uint `gorm:"column:land_details_id;primaryKey;autoIncrement"`
	PropertyID    uint `gorm:"column:property_id;not null"`

	Rtc        string `gorm:"column:rtc;type:varchar(10)"`
	Ec         string `gorm:"column:ec;type:varchar(10)"`
	SyNo       string `gorm:"column:sy_no;type:varchar(100)"`
	KhataNo    string `gorm:"column:khata_no;type:varchar(100)"`
	MrNo       string `gorm:"column:mr_no;type:varchar(100)"`
	Acre       string `gorm:"column:acre;type:varchar(50)"`
	Gunte      string `gorm:"column:gunte;type:varchar(50)"`
	Karab      string `gorm:"column:karab;type:varchar(50)"`
	KhuskiTari string `gorm:"column:khuski_tari;type:varchar(50)"`
	Converted  string `gorm:"column:converted;type:varchar(10)"`
	Purpose    string `gorm:"column:purpose;type:varchar(100)"`

	Property Property `gorm:"foreignKey:PropertyID;references:PropertyID"`
}

/* =========================
   property_tax_details
========================= */

type PropertyTaxDetails struct {
	TaxID      uint `gorm:"column:tax_id;primaryKey;autoIncrement"`
	PropertyID uint `gorm:"column:property_id;not null"`

	TaxYear                      string  `gorm:"column:tax_year;type:varchar(10)"`
	TaxPaid                      string  `gorm:"column:tax_paid;type:varchar(10)"`
	TaxAmount                    float64 `gorm:"column:tax_amount"`
	ReceiptNo                    string  `gorm:"column:receipt_no;type:varchar(100)"`
	TaxAmountPaidPreviousYear    float64 `gorm:"column:tax_amount_paid_previous_year"`
	TaxAmountPaidCurrentYear     float64 `gorm:"column:tax_amount_paid_current_year"`
	ReceiptPhotoLink             string  `gorm:"column:receipt_photo_link;type:text"`
	PTax                         string  `gorm:"column:p_tax;type:varchar(20)"`
	ReceivedFrom                 string  `gorm:"column:received_from;type:varchar(100)"`
	AcquisitionType              string  `gorm:"column:acquisition_type;type:varchar(50)"`

	Property Property `gorm:"foreignKey:PropertyID;references:PropertyID"`
}

/* =========================
   property_ownership_details
========================= */

type PropertyOwnershipDetails struct {
	OwnershipID uint `gorm:"column:ownership_id;primaryKey;autoIncrement"`
	PropertyID  uint `gorm:"column:property_id;not null"`

	ReceivedFrom        string `gorm:"column:received_from;type:varchar(100)"`
	AcquisitionType     string `gorm:"column:acquisition_type;type:varchar(50)"`
	RegistrationDetails string `gorm:"column:registration_details;type:text"`
	Title               string `gorm:"column:title;type:varchar(100)"`
	Incharge            string `gorm:"column:incharge;type:varchar(100)"`
	PhoneNumber         string `gorm:"column:phone_number;type:varchar(15)"`

	Property Property `gorm:"foreignKey:PropertyID;references:PropertyID"`
}

/* =========================
   property_documents
========================= */

type PropertyDocuments struct {
	DocumentID uint `gorm:"column:document_id;primaryKey;autoIncrement"`
	PropertyID uint `gorm:"column:property_id;not null"`

	DocumentType      string `gorm:"column:document_type;type:varchar(100)"`
	DocumentLink      string `gorm:"column:document_link;type:text"`
	OriginalAvailable string `gorm:"column:original_available;type:varchar(10)"`

	Property Property `gorm:"foreignKey:PropertyID;references:PropertyID"`
}

/* =========================
   property_building_details
========================= */

type PropertyBuildingDetails struct {
	BuildingDetailsID uint `gorm:"column:building_details_id;primaryKey;autoIncrement"`
	PropertyID        uint `gorm:"column:property_id;not null"`

	PlotSize           string `gorm:"column:plot_size;type:varchar(50)"`
	BuiltUpArea        string `gorm:"column:built_up_area;type:varchar(50)"`
	YearOfConstruction string `gorm:"column:year_of_construction;type:varchar(10)"`
	ApplicationNo      string `gorm:"column:application_no;type:varchar(100)"`

	Property Property `gorm:"foreignKey:PropertyID;references:PropertyID"`
}

/* =========================
   property_media
========================= */

type PropertyMedia struct {
	MediaID         uint `gorm:"column:media_id;primaryKey;autoIncrement"`
	PropertyID      uint `gorm:"column:property_id;not null"`
	ScannedDeedLink string `gorm:"column:scanned_deed_link;type:text"`
	PhotoLink       string `gorm:"column:photo_link;type:text"`
	VideoLink       string `gorm:"column:video_link;type:text"`
	Remarks         string `gorm:"column:remarks;type:text"`

	Property Property `gorm:"foreignKey:PropertyID;references:PropertyID"`
}
