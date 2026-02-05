package routes

import (
	"github.com/gin-gonic/gin"
	"property-backend/controllers"
)

// AssetRoutes registers asset endpoints under /assets
func AssetRoutes(rg *gin.RouterGroup, controller *controllers.AssetController) {
	assets := rg.Group("/assets")
	{
		// Add asset
		// @Summary Create an asset
		// @Tags Assets
		// @Accept json
		// @Produce json
		// @Router /api/v1/assets [post]
		assets.POST("", controller.AddAsset)

		// Get all assets
		// @Summary Get all assets
		// @Tags Assets
		// @Produce json
		// @Router /api/v1/assets [get]
		assets.GET("", controller.GetAllAssets)
	}
}
