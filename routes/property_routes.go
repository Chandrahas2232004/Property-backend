package routes

import (
	"github.com/gin-gonic/gin"
	"property-backend/controllers"
)

// PropertyRoutes registers property-related routes under /properties
func PropertyRoutes(rg *gin.RouterGroup, controller *controllers.PropertyController) {
	props := rg.Group("/properties")
	{
		// Total properties
		// @Summary Get total properties count
		// @Tags Properties
		// @Produce json
		// @Router /api/v1/properties/total [get]
		props.GET("/total", controller.TotalProperties)

		// Active rental properties count
		// @Summary Get active rental properties count
		// @Tags Properties
		// @Produce json
		// @Router /api/v1/properties/active-rental/count [get]
		props.GET("/active-rental/count", controller.ActiveRentalPropertyCount)

		// Add property
		// @Summary Add a new property
		// @Tags Properties
		// @Accept json
		// @Produce json
		// @Router /api/v1/properties [post]
		props.POST("", controller.AddProperty)

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
	}
}
