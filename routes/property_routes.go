package routes

import (
	"property-backend/controllers"

	"github.com/gin-gonic/gin"
)

// PropertyRoutes registers property-related routes under /properties
func PropertyRoutes(rg *gin.RouterGroup, controller *controllers.PropertyController) {
	props := rg.Group("/properties")
	{
		// Add property basic info (Step 1: Text data only)
		// @Summary Add property basic information
		// @Tags Properties
		// @Accept json
		// @Produce json
		// @Router /api/v1/properties/basic [post]
		props.POST("/basic", controller.AddPropertyBasicInfo)

		// Upload property files (Step 2: Binary files)
		// @Summary Upload property files
		// @Tags Properties
		// @Accept multipart/form-data
		// @Produce json
		// @Router /api/v1/properties/:id/files [post]
		props.POST("/:id/files", controller.UploadPropertyFiles)

		// Agricultural properties
		// @Summary List agricultural land properties
		// @Tags Properties
		// @Produce json
		// @Router /api/v1/properties/agricultural [get]
		props.GET("/agricultural", controller.AgriculturalLandProperties)

		// Residential properties
		// @Summary List residential land properties
		// @Tags Properties
		// @Produce json
		// @Router /api/v1/properties/residential [get]
		props.GET("/residential", controller.ResidentialLandProperties)

		// Commercial properties
		// @Summary List commercial land properties
		// @Tags Properties
		// @Produce json
		// @Router /api/v1/properties/commercial [get]
		props.GET("/commercial", controller.CommercialLandProperties)

		// Get all properties
		// @Summary Get all properties
		// @Description Retrieve all properties with complete details
		// @Tags Properties
		// @Produce json
		// @Router /api/v1/properties [get]
		props.GET("", controller.GetAllProperties)
	}
}
