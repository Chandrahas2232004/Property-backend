package models

import "time"

/* =========================
   country_master
========================= */

type CountryMaster struct {
	CountryID   uint   `gorm:"column:country_id;primaryKey;autoIncrement" json:"country_id"`
	CountryName string `gorm:"column:country_name;type:varchar(100);not null;unique" json:"country_name"`
}

/* =========================
   state_master
========================= */

type StateMaster struct {
	StateID   uint   `gorm:"column:state_id;primaryKey;autoIncrement" json:"state_id"`
	StateName string `gorm:"column:state_name;type:varchar(100);not null" json:"state_name"`
	CountryID uint   `gorm:"column:country_id;not null" json:"country_id"`

	Country CountryMaster `gorm:"foreignKey:CountryID;references:CountryID"`
}

/* =========================
   district_master
========================= */

type DistrictMaster struct {
	DistrictID   uint   `gorm:"column:district_id;primaryKey;autoIncrement" json:"district_id"`
	DistrictName string `gorm:"column:district_name;type:varchar(100);not null" json:"district_name"`
	StateID      uint   `gorm:"column:state_id;not null" json:"state_id"`

	State StateMaster `gorm:"foreignKey:StateID;references:StateID"`
}

/* =========================
   taluk_master
========================= */

type TalukMaster struct {
	TalukID    uint   `gorm:"column:taluk_id;primaryKey;autoIncrement" json:"taluk_id"`
	TalukName  string `gorm:"column:taluk_name;type:varchar(100);not null" json:"taluk_name"`
	DistrictID uint   `gorm:"column:district_id;not null" json:"district_id"`

	District DistrictMaster `gorm:"foreignKey:DistrictID;references:DistrictID"`
}

/* =========================
   address
========================= */

type Address struct {
	AddressID      uint   `gorm:"column:address_id;primaryKey;autoIncrement" json:"address_id"`
	TalukID        uint   `gorm:"column:taluk_id;not null" json:"taluk_id"`
	Hobli          string `gorm:"column:hobli;type:varchar(100)" json:"hobli"`
	Village        string `gorm:"column:village;type:varchar(100)" json:"village"`
	StreetAddress  string `gorm:"column:street_address;type:text" json:"street_address"`
	Pincode        string `gorm:"column:pincode;type:varchar(10)" json:"pincode"`
	LatCoordinate  string `gorm:"column:lat_coordinate;type:varchar(50)" json:"lat_coordinate"`
	LongCoordinate string `gorm:"column:long_coordinate;type:varchar(50)" json:"long_coordinate"`

	Taluk TalukMaster `gorm:"foreignKey:TalukID;references:TalukID"`
}

/* =========================
   property
========================= */

type Property struct {
	PropertyID     uint      `gorm:"column:property_id;primaryKey;autoIncrement" json:"property_id"`
	PropertyName   string    `gorm:"column:property_name;type:varchar(150);not null" json:"property_name"`
	PropertyTypeID uint      `gorm:"column:property_type_id;not null" json:"property_type_id"`
	AddressID      uint      `gorm:"column:address_id;not null" json:"address_id"`
	UserID         uint      `gorm:"column:user_id;not null" json:"user_id"`
	Value          string    `gorm:"column:value;type:varchar(50)" json:"value"`
	Income         string    `gorm:"column:income;type:varchar(50)" json:"income"`
	OriginalDeed   string    `gorm:"column:original_deed;type:varchar(10)" json:"original_deed"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`

	Address      Address            `gorm:"foreignKey:AddressID;references:AddressID;constraint:OnDelete:RESTRICT"`
	PropertyType PropertyTypeMaster `gorm:"foreignKey:PropertyTypeID;references:PropertyTypeID"`

	LandDetails     *PropertyLandDetails      `gorm:"foreignKey:PropertyID"`
	TaxDetails      *PropertyTaxDetails       `gorm:"foreignKey:PropertyID"`
	Ownership       *PropertyOwnershipDetails `gorm:"foreignKey:PropertyID"`
	BuildingDetails *PropertyBuildingDetails  `gorm:"foreignKey:PropertyID"`
	Media           *PropertyMedia            `gorm:"foreignKey:PropertyID"`
}

func (Property) TableName() string {
	return "property"
}

/* =========================
   property_type_master
========================= */

type PropertyTypeMaster struct {
	PropertyTypeID   uint   `gorm:"column:property_type_id;primaryKey;autoIncrement" json:"property_type_id"`
	PropertyTypeName string `gorm:"column:property_type_name;type:varchar(100);not null;unique" json:"property_type_name"`
}

func (PropertyTypeMaster) TableName() string {
	return "property_type_master"
}

/* =========================
   property_land_details (1–1)
========================= */

type PropertyLandDetails struct {
	LandDetailsID uint `gorm:"column:land_details_id;primaryKey;autoIncrement" json:"land_details_id"`
	PropertyID    uint `gorm:"column:property_id;not null;unique" json:"property_id"`

	Rtc       string `gorm:"column:rtc;type:varchar(10)" json:"rtc"`
	Ec        string `gorm:"column:ec;type:varchar(10)" json:"ec"`
	SyNo      string `gorm:"column:sy_no;type:varchar(100)" json:"sy_no"`
	KhataNo   string `gorm:"column:khata_no;type:varchar(100)" json:"khata_no"`
	MrNo      string `gorm:"column:mr_no;type:varchar(100)" json:"mr_no"`
	Acre      string `gorm:"column:acre;type:varchar(50)" json:"acre"`
	Gunte     string `gorm:"column:gunte;type:varchar(50)" json:"gunte"`
	Karab     string `gorm:"column:karab;type:varchar(50)" json:"karab"`
	Converted string `gorm:"column:converted;type:varchar(10)" json:"converted"`
	Purpose   string `gorm:"column:purpose;type:varchar(100)" json:"purpose"`
	KhuskiTari string `gorm:"column:khuski_tari;type:varchar(50)" json:"khuskitari"`

	Property *Property `gorm:"foreignKey:PropertyID;references:PropertyID;constraint:OnDelete:CASCADE"`
}

/* =========================
   property_tax_details (1–1)
========================= */

type PropertyTaxDetails struct {
	TaxID      uint `gorm:"column:tax_id;primaryKey;autoIncrement" json:"tax_id"`
	PropertyID uint `gorm:"column:property_id;not null;unique" json:"property_id"`

	TaxPaid      bool    `gorm:"column:tax_paid" json:"tax_paid"`
	ReceiptNo   string  `gorm:"column:receipt_no;type:varchar(100)" json:"receipt_no"`
	PrevAmount  float64 `gorm:"column:tax_amount_paid_previous_year" json:"tax_amount_paid_previous_year"`
	CurrAmount  float64 `gorm:"column:tax_amount_paid_current_year" json:"tax_amount_paid_current_year"`
	ReceiptLink string  `gorm:"column:receipt_photo_link;type:text" json:"receipt_photo_link"`

	Property *Property `gorm:"foreignKey:PropertyID;references:PropertyID;constraint:OnDelete:CASCADE"`
}

/* =========================
   property_ownership_details (1–1)
========================= */

type PropertyOwnershipDetails struct {
	OwnershipID uint `gorm:"column:ownership_id;primaryKey;autoIncrement" json:"ownership_id"`
	PropertyID  uint `gorm:"column:property_id;not null;unique" json:"property_id"`

	ReceivedFrom        string `gorm:"column:received_from;type:varchar(100)" json:"received_from"`
	AcquisitionType     string `gorm:"column:acquisition_type;type:varchar(50)" json:"acquisition_type"`
	RegistrationDetails string `gorm:"column:registration_details;type:text" json:"registration_details"`
	Title               string `gorm:"column:title;type:varchar(100)" json:"title"`
	Incharge            string `gorm:"column:incharge;type:varchar(100)" json:"incharge"`
	PhoneNumber         string `gorm:"column:phone_number;type:varchar(15)" json:"phone_number"`

	Property *Property `gorm:"foreignKey:PropertyID;references:PropertyID;constraint:OnDelete:CASCADE"`
}

/* =========================
   property_building_details (1–1)
========================= */

type PropertyBuildingDetails struct {
	BuildingDetailsID uint `gorm:"column:building_details_id;primaryKey;autoIncrement" json:"building_details_id"`
	PropertyID        uint `gorm:"column:property_id;not null;unique" json:"property_id"`

	PlotSize           string `gorm:"column:plot_size;type:varchar(50)" json:"plot_size"`
	BuiltUpArea        string `gorm:"column:built_up_area;type:varchar(50)" json:"built_up_area"`
	YearOfConstruction string `gorm:"column:year_of_construction;type:varchar(10)" json:"year_of_construction"`
	ApplicationNo      string `gorm:"column:application_no;type:varchar(100)" json:"application_no"`

	Property *Property `gorm:"foreignKey:PropertyID;references:PropertyID;constraint:OnDelete:CASCADE"`
}

/* =========================
   property_media (1–1)
========================= */

type PropertyMedia struct {
	MediaID    uint `gorm:"column:media_id;primaryKey;autoIncrement" json:"media_id"`
	PropertyID uint `gorm:"column:property_id;not null;unique" json:"property_id"`

	ScannedDeedLink string `gorm:"column:scanned_deed_link;type:text" json:"scanned_deed_link"`
	PhotoLink       string `gorm:"column:photo_link;type:text" json:"photo_link"`
	Remarks         string `gorm:"column:remarks;type:text" json:"remarks"`

	Property *Property `gorm:"foreignKey:PropertyID;references:PropertyID;constraint:OnDelete:CASCADE"`
}
