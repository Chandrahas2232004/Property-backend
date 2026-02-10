package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"property-backend/services"
)

// FormController handles form-related endpoints
type FormController struct {
	svc services.FormService
}

// NewFormController creates a new FormController
func NewFormController(svc services.FormService) *FormController {
	return &FormController{svc: svc}
}

// GetCountry retrieves all countries
// GetCountry godoc
// @Summary Get all countries
// @Tags Form
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/form/country [get]
func (fc *FormController) GetCountry(c *gin.Context) {
	result, err := fc.svc.GetCountry(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// GetState retrieves all states by country ID
// GetState godoc
// @Summary Get states by country ID
// @Tags Form
// @Produce json
// @Param id path string true "Country ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/form/state/{id} [get]
func (fc *FormController) GetState(c *gin.Context) {
	id := c.Param("id")

	result, err := fc.svc.GetState(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// GetDistrict retrieves all districts by state ID
// GetDistrict godoc
// @Summary Get districts by state ID
// @Tags Form
// @Produce json
// @Param id path string true "State ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/form/district/{id} [get]
func (fc *FormController) GetDistrict(c *gin.Context) {
	id := c.Param("id")

	result, err := fc.svc.GetDistrict(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// GetTaluk retrieves all taluks by district ID
// GetTaluk godoc
// @Summary Get taluks by district ID
// @Tags Form
// @Produce json
// @Param id path string true "District ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/form/taluk/{id} [get]
func (fc *FormController) GetTaluk(c *gin.Context) {
	id := c.Param("id")

	result, err := fc.svc.GetTaluk(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
