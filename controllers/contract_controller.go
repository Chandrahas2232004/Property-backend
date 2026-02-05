package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"property-backend/models"
	"property-backend/services"
)

// ContractController handles contract endpoints
type ContractController struct {
	svc services.ContractService
}

// NewContractController creates a new ContractController
func NewContractController(svc services.ContractService) *ContractController {
	return &ContractController{svc: svc}
}

// AddContract godoc
// @Summary Create a contract
// @Tags Contracts
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/v1/contracts [post]
func (ctr *ContractController) AddContract(c *gin.Context) {
	var req struct {
		Name           string  `json:"name" binding:"required"`
		ContractTypeID uint    `json:"contract_type_id" binding:"required"`
		RelatedTo      string  `json:"related_to"`
		Cost           float64 `json:"cost"`
		Provider       string  `json:"provider"`
		StartDate      string  `json:"start_date"`
		EndDate        string  `json:"end_date"`
		Terms          string  `json:"terms"`
		AssetID        uint    `json:"asset_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	co := models.Contract{
		Name:           req.Name,
		ContractTypeID: req.ContractTypeID,
		RelatedTo:      req.RelatedTo,
		Cost:           req.Cost,
		Provider:       req.Provider,
		Terms:          req.Terms,
		AssetID:        req.AssetID,
	}
	id, err := ctr.svc.AddContract(context.Background(), &co)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetAllContracts godoc
// @Summary Get all contracts
// @Tags Contracts
// @Produce json
// @Success 200 {array} models.Contract
// @Router /api/v1/contracts [get]
func (ctr *ContractController) GetAllContracts(c *gin.Context) {
	contracts, err := ctr.svc.GetAllContracts(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contracts)
}

// GetLeaseContracts godoc
// @Summary Get all lease contracts
// @Tags Contracts
// @Produce json
// @Success 200 {array} models.Contract
// @Router /api/v1/contracts/lease [get]
func (ctr *ContractController) GetLeaseContracts(c *gin.Context) {
	contracts, err := ctr.svc.GetContractsByType(context.Background(), "lease")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contracts)
}

// GetAMCContracts godoc
// @Summary Get all AMC contracts
// @Tags Contracts
// @Produce json
// @Success 200 {array} models.Contract
// @Router /api/v1/contracts/amc [get]
func (ctr *ContractController) GetAMCContracts(c *gin.Context) {
	contracts, err := ctr.svc.GetContractsByType(context.Background(), "amc")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contracts)
}
