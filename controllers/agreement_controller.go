package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"property-backend/models"
	"property-backend/services"
)

// AgreementController handles agreement endpoints
type AgreementController struct {
	svc services.AgreementService
}

// NewAgreementController creates a new AgreementController
func NewAgreementController(svc services.AgreementService) *AgreementController {
	return &AgreementController{svc: svc}
}

// AddRentalAgreement godoc
// @Summary Create a rental agreement
// @Tags Agreements
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/v1/agreements [post]
func (a *AgreementController) AddRentalAgreement(c *gin.Context) {
	var req struct {
		PropertyID uint    `json:"property_id" binding:"required"`
		TenantName string  `json:"tenant_name" binding:"required"`
		ContactNo  string  `json:"contact_no"`
		Rent       float64 `json:"rent" binding:"required"`
		Deposit    float64 `json:"deposit"`
		StartDate  string  `json:"start_date" binding:"required"`
		EndDate    string  `json:"end_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ag := models.Agreement{
		PropertyID: req.PropertyID,
		TenantName: req.TenantName,
		ContactNo:  req.ContactNo,
		Rent:       req.Rent,
		Deposit:    req.Deposit,
	}
	// parse dates into time.Time when repository expects them (gorm will handle zero values)
	id, err := a.svc.AddRentalAgreement(context.Background(), &ag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetAllAgreements godoc
// @Summary Get all agreements
// @Tags Agreements
// @Produce json
// @Success 200 {array} models.Agreement
// @Router /api/v1/agreements [get]
func (a *AgreementController) GetAllAgreements(c *gin.Context) {
	ags, err := a.svc.GetAllAgreements(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ags)
}
