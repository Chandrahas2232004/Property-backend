package routes

import (
	"property-backend/controllers"

	"github.com/gin-gonic/gin"
)

// DashboardRoutes registers dashboard endpoints under /dashboard
func DashboardRoutes(rg *gin.RouterGroup, controller *controllers.DashboardController) {
	dashboard := rg.Group("/dashboard")
	{
		// Total properties
		// @Summary Get total property count
		// @Tags Dashboard
		// @Produce json
		// @Router /api/v1/dashboard/total-properties [get]
		dashboard.GET("/total-properties", controller.GetTotalPropertyCount)

		// Total assets
		// @Summary Get total asset count
		// @Tags Dashboard
		// @Produce json
		// @Router /api/v1/dashboard/total-assets [get]
		dashboard.GET("/total-assets", controller.GetTotalAssetCount)

		// Total AMC contracts
		// @Summary Get total AMC contract count
		// @Tags Dashboard
		// @Produce json
		// @Router /api/v1/dashboard/total-amc-contracts [get]
		dashboard.GET("/total-amc-contracts", controller.GetTotalAMCContractCount)

		// Total active rental properties
		// @Summary Get total active rental count
		// @Tags Dashboard
		// @Produce json
		// @Router /api/v1/dashboard/active-rental-count [get]
		dashboard.GET("/active-rental-count", controller.GetTotalActiveRentalCount)

		// Recent activities
		// @Summary Get recent activities
		// @Tags Dashboard
		// @Produce json
		// @Router /api/v1/dashboard/recent-activities [get]
		dashboard.GET("/recent-activities", controller.GetRecentActivities)
	}
}
