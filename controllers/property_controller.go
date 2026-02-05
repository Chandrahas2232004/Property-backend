package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"property-backend/services"
)

// PropertyController handles property-related endpoints
type PropertyController struct {
	svc services.PropertyService
}

// NewPropertyController creates a new PropertyController
func NewPropertyController(svc services.PropertyService) *PropertyController {
	return &PropertyController{svc: svc}
}

// TotalProperties godoc
// @Summary Get total properties count
// @Tags Properties
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/properties/total [get]
func (p *PropertyController) TotalProperties(c *gin.Context) {
	total, err := p.svc.Total(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": total})
}

// ActiveRentalPropertyCount godoc
// @Summary Get active rental properties count
// @Tags Properties
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/properties/active-rental/count [get]
func (p *PropertyController) ActiveRentalPropertyCount(c *gin.Context) {
	count, err := p.svc.ActiveRentalCount(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"active_rental_count": count})
}

// AddProperty godoc
// @Summary Add a new property with all details
// @Tags Properties
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/v1/properties [post]
func (p *PropertyController) AddProperty(c *gin.Context) {
	var req struct {
		// Property basic info
		PropertyName   string `json:"property_name" binding:"required"`
		PropertyTypeID uint   `json:"property_type_id"`
		PropertyTypeName string `json:"property_type_name"`
		Value          string `json:"value"`
		Income         string `json:"income"`
		OriginalDeed   string `json:"original_deed"`
		UserID         uint   `json:"user_id" binding:"required"`

		// Address information
		CountryName    string `json:"country_name"`
		StateName      string `json:"state_name"`
		DistrictName   string `json:"district_name"`
		TalukName      string `json:"taluk_name"`
		Hobli          string `json:"hobli"`
		Village        string `json:"village"`
		StreetAddress  string `json:"street_address"`
		Pincode        string `json:"pincode"`
		LatCoordinate  string `json:"lat_coordinate"`
		LongCoordinate string `json:"long_coordinate"`

		// Land details
		Rtc        string `json:"rtc"`
		Ec         string `json:"ec"`
		SyNo       string `json:"syno"`
		KhataNo    string `json:"khatano"`
		MrNo       string `json:"mrno"`
		Acre       string `json:"acre"`
		Gunte      string `json:"gunte"`
		Karab      string `json:"karab"`
		Converted  string `json:"converted"`
		Purpose    string `json:"purpose"`
		KhuskiTari string `json:"khuskitari"`

		// Tax details
		ReceiptNo   string  `json:"receipt_no"`
		PrevAmount  float64 `json:"prev_amount"`
		CurrAmount  float64 `json:"curr_amount"`
		ReceiptLink string  `json:"receipt_link"`

		// Ownership details
		ReceivedFrom        string `json:"received_from"`
		AcquisitionType     string `json:"acquisition_type"`
		RegistrationDetails string `json:"registration_details"`
		Title               string `json:"title"`
		Incharge            string `json:"incharge"`
		PhoneNumber         string `json:"phone_number"`

		// Building details
		PlotSize           string `json:"plot_size"`
		BuiltUpArea        string `json:"built_up_area"`
		YearOfConstruction string `json:"year_of_construction"`
		ApplicationNo      string `json:"application_no"`

		// Media
		ScannedDeedLink string `json:"scanned_deed_link"`
		PhotoLink       string `json:"photo_link"`
		Remarks         string `json:"remarks"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert struct to map for service processing
	reqBytes, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process request"})
		return
	}

	var reqMap map[string]interface{}
	if err := json.Unmarshal(reqBytes, &reqMap); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process request"})
		return
	}

	id, err := p.svc.AddProperty(context.Background(), reqMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// AgriculturalLandProperties godoc
// @Summary List agricultural land properties
// @Tags Properties
// @Produce json
// @Success 200 {array} models.Property
// @Router /api/v1/properties/agricultural [get]
func (p *PropertyController) AgriculturalLandProperties(c *gin.Context) {
	props, err := p.svc.ListByType(context.Background(), "agricultural")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, props)
}

// ResidentialLandProperties godoc
// @Summary List residential land properties
// @Tags Properties
// @Produce json
// @Success 200 {array} models.Property
// @Router /api/v1/properties/residential [get]
func (p *PropertyController) ResidentialLandProperties(c *gin.Context) {
	props, err := p.svc.ListByType(context.Background(), "residential")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, props)
}

// CommercialLandProperties godoc
// @Summary List commercial land properties
// @Tags Properties
// @Produce json
// @Success 200 {array} models.Property
// @Router /api/v1/properties/commercial [get]
func (p *PropertyController) CommercialLandProperties(c *gin.Context) {
	props, err := p.svc.ListByType(context.Background(), "commercial")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, props)
}
