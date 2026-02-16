package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"property-backend/models"
	"property-backend/services"

	"github.com/gin-gonic/gin"
)

// PropertyController handles property-related endpoints
type PropertyController struct {
	svc services.PropertyService
}

// PropertyListResponse documents the GetAllProperties response shape.
type PropertyListResponse struct {
	Count int               `json:"count"`
	Data  []models.Property `json:"data"`
}

// NewPropertyController creates a new PropertyController
func NewPropertyController(svc services.PropertyService) *PropertyController {
	return &PropertyController{svc: svc}
}

// AddPropertyBasicInfo godoc
// @Summary Add a new property with text information only (Step 1)
// @Description Creates property with all text fields. Use UploadPropertyFiles endpoint to upload binary files after.
// @Tags Properties
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/v1/properties/basic [post]
func (p *PropertyController) AddPropertyBasicInfo(c *gin.Context) {
	var req struct {
		// Property basic info
		PropertyName     string `json:"property_name" binding:"required"`
		PropertyTypeID   uint   `json:"property_type_id"`
		PropertyTypeName string `json:"property_type_name"`
		Value            string `json:"value"`
		Income           string `json:"income"`
		OriginalDeed     string `json:"original_deed"`
		UserID           uint   `json:"user_id" binding:"required"`

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
		ReceiptNo  string  `json:"receipt_no"`
		PrevAmount float64 `json:"prev_amount"`
		CurrAmount float64 `json:"curr_amount"`

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

		// Media (text only)
		PhotoLink string `json:"photo_link"`
		Remarks   string `json:"remarks"`
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

	id, err := p.svc.AddPropertyBasicInfo(context.Background(), reqMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"property_id": id,
		"message":     "Property basic info created successfully. Upload files using /properties/:id/files endpoint",
	})
}

// UploadPropertyFiles godoc
// @Summary Upload binary files for a property (Step 2)
// @Description Upload ScannedDeed and Receipt files to S3 for an existing property
// @Tags Properties
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Property ID"
// @Param scanned_deed formData file false "Scanned Deed PDF/Image"
// @Param receipt formData file false "Tax Receipt PDF/Image"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/properties/{id}/files [post]
func (p *PropertyController) UploadPropertyFiles(c *gin.Context) {
	propertyIDStr := c.Param("id")
	propertyID, err := strconv.ParseUint(propertyIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid property_id"})
		return
	}

	filesData := map[string]interface{}{
		"property_id": uint(propertyID),
	}

	// Handle scanned_deed file
	// Note: Do NOT defer close here - file must remain open until after UploadPropertyFiles completes
	scannedDeedFile, scannedDeedHeader, err := c.Request.FormFile("scanned_deed")
	if err == nil {
		filesData["scanned_deed_file"] = scannedDeedFile
		filesData["scanned_deed_header"] = scannedDeedHeader
	}

	// Handle receipt file
	// Note: Do NOT defer close here - file must remain open until after UploadPropertyFiles completes
	receiptFile, receiptHeader, err := c.Request.FormFile("receipt")
	if err == nil {
		filesData["receipt_file"] = receiptFile
		filesData["receipt_header"] = receiptHeader
	}

	// If no files provided
	if filesData["scanned_deed_file"] == nil && filesData["receipt_file"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "at least one file must be provided"})
		return
	}

	// Upload files to S3 and store URLs in database
	uploadedFiles, err := p.svc.UploadPropertyFiles(context.Background(), filesData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Files uploaded successfully to S3",
		"property_id":    propertyID,
		"uploaded_files": uploadedFiles,
	})
}

// AgriculturalLandProperties godoc
// @Summary List agricultural land properties for the user
// @Tags Properties
// @Produce json
// @Success 200 {array} models.Property
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/properties/agricultural [get]
func (p *PropertyController) AgriculturalLandProperties(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	props, err := p.svc.ListByTypeByUser(context.Background(), "agricultural", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, props)
}

// ResidentialLandProperties godoc
// @Summary List residential land properties for the user
// @Tags Properties
// @Produce json
// @Success 200 {array} models.Property
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/properties/residential [get]
func (p *PropertyController) ResidentialLandProperties(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	props, err := p.svc.ListByTypeByUser(context.Background(), "residential", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, props)
}

// CommercialLandProperties godoc
// @Summary List commercial land properties for the user
// @Tags Properties
// @Produce json
// @Success 200 {array} models.Property
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/properties/commercial [get]
func (p *PropertyController) CommercialLandProperties(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	props, err := p.svc.ListByTypeByUser(context.Background(), "commercial", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, props)
}

// GetAllProperties godoc
// @Summary Get all properties for the user
// @Description Retrieve all properties for the authenticated user with complete details including address, land details, tax details, media, and ownership information
// @Tags Properties
// @Produce json
// @Success 200 {object} PropertyListResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/properties [get]
func (p *PropertyController) GetAllProperties(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	props, err := p.svc.GetAllByUser(context.Background(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Handle empty result
	if len(props) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"count": 0,
			"data":  []interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(props),
		"data":  props,
	})
}
