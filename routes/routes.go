package routes

import (
	"github.com/gin-gonic/gin"
	"property-backend/controllers"
)

// RegisterRoutes registers all API routes under a provided RouterGroup
func RegisterRoutes(api *gin.RouterGroup,
	authController *controllers.AuthController,
	propertyController *controllers.PropertyController,
	agreementController *controllers.AgreementController,
	assetController *controllers.AssetController,
	contractController *controllers.ContractController,
) {
	// domain-specific routes
	AuthRoutes(api, authController)
	PropertyRoutes(api, propertyController)
	AgreementRoutes(api, agreementController)
	AssetRoutes(api, assetController)
	ContractRoutes(api, contractController)
}
