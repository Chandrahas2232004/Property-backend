package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"property-backend/models"
	"property-backend/services"
)

// AssetController handles asset endpoints
type AssetController struct {
	svc services.AssetService
}

// NewAssetController creates a new AssetController
func NewAssetController(svc services.AssetService) *AssetController {
	return &AssetController{svc: svc}
}

// AddAsset godoc
// @Summary Create an asset
// @Tags Assets
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/v1/assets [post]
func (a *AssetController) AddAsset(c *gin.Context) {
	var req struct {
		AssetName        string `json:"asset_name" binding:"required"`
		AssetTypeID      uint   `json:"asset_type_id" binding:"required"`
		PropertyID       uint   `json:"property_id" binding:"required"`
		AssetLocation    string `json:"location"`
		AssetCost        string `json:"cost"`
		AssetAMCProvider string `json:"amc_provider"`
		AssetStartDate   string `json:"start_date"`
		AssetEndDate     string `json:"end_date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	asset := models.Asset{
		AssetName:        req.AssetName,
		AssetTypeID:      req.AssetTypeID,
		PropertyID:       req.PropertyID,
		AssetLocation:    req.AssetLocation,
		AssetCost:        req.AssetCost,
		AssetAMCProvider: req.AssetAMCProvider,
		AssetStartDate:   req.AssetStartDate,
		AssetEndDate:     req.AssetEndDate,
	}
	id, err := a.svc.AddAsset(context.Background(), &asset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetAllAssets godoc
// @Summary Get all assets
// @Tags Assets
// @Produce json
// @Success 200 {array} models.Asset
// @Router /api/v1/assets [get]
func (a *AssetController) GetAllAssets(c *gin.Context) {
	assets, err := a.svc.GetAllAssets(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, assets)
}
