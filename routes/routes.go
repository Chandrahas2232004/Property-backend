package routes

import (
	"property-backend/controllers"
	"property-backend/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API routes under a provided RouterGroup
func RegisterRoutes(api *gin.RouterGroup,
	authController *controllers.AuthController,
	propertyController *controllers.PropertyController,
	agreementController *controllers.AgreementController,
	assetController *controllers.AssetController,
	contractController *controllers.ContractController,
	formController *controllers.FormController,
	dashboardController *controllers.DashboardController,
) {
	// Public routes (no auth required)
	AuthRoutes(api, authController)

	// Protected routes (auth required)
	protectedAPI := api.Group("")
	protectedAPI.Use(middleware.AuthMiddleware())
	{
		PropertyRoutes(protectedAPI, propertyController)
		AgreementRoutes(protectedAPI, agreementController)
		AssetRoutes(protectedAPI, assetController)
		ContractRoutes(protectedAPI, contractController)
		FormRoutes(protectedAPI, formController)
		DashboardRoutes(protectedAPI, dashboardController)
	}
}
