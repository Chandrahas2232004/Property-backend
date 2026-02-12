package controllers

import (
	"net/http"

	"property-backend/services"

	"github.com/gin-gonic/gin"
)

// DashboardController handles dashboard endpoints
// It aggregates high-level metrics for the UI.
type DashboardController struct {
	svc services.DashboardService
}

// NewDashboardController creates a new DashboardController
func NewDashboardController(svc services.DashboardService) *DashboardController {
	return &DashboardController{svc: svc}
}

// GetTotalPropertyCount godoc
// @Summary Get total property count
// @Tags Dashboard
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/dashboard/total-properties [get]
func (d *DashboardController) GetTotalPropertyCount(c *gin.Context) {
	count, err := d.svc.TotalPropertyCount(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_property_count": count})
}

// GetTotalAssetCount godoc
// @Summary Get total asset count
// @Tags Dashboard
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/dashboard/total-assets [get]
func (d *DashboardController) GetTotalAssetCount(c *gin.Context) {
	count, err := d.svc.TotalAssetCount(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_asset_count": count})
}

// GetTotalAMCContractCount godoc
// @Summary Get total AMC contract count
// @Tags Dashboard
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/dashboard/total-amc-contracts [get]
func (d *DashboardController) GetTotalAMCContractCount(c *gin.Context) {
	count, err := d.svc.TotalAMCContractCount(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_amc_contract_count": count})
}

// GetTotalActiveRentalCount godoc
// @Summary Get total active rental count
// @Tags Dashboard
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/dashboard/active-rental-count [get]
func (d *DashboardController) GetTotalActiveRentalCount(c *gin.Context) {
	count, err := d.svc.TotalActiveRentalCount(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_active_rental_count": count})
}
