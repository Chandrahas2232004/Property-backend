package routes

import (
	"property-backend/controllers"

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
	// domain-specific routes
	AuthRoutes(api, authController)
	PropertyRoutes(api, propertyController)
	AgreementRoutes(api, agreementController)
	AssetRoutes(api, assetController)
	ContractRoutes(api, contractController)
	FormRoutes(api, formController)
	DashboardRoutes(api, dashboardController)
}
