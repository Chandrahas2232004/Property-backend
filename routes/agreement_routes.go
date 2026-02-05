package routes

import (
	"github.com/gin-gonic/gin"
	"property-backend/controllers"
)

// AgreementRoutes registers agreement endpoints under /agreements
func AgreementRoutes(rg *gin.RouterGroup, controller *controllers.AgreementController) {
	agmts := rg.Group("/agreements")
	{
		// Add rental agreement
		// @Summary Create a rental agreement
		// @Tags Agreements
		// @Accept json
		// @Produce json
		// @Router /api/v1/agreements [post]
		agmts.POST("", controller.AddRentalAgreement)

		// Get all agreements
		// @Summary Get all agreements
		// @Tags Agreements
		// @Produce json
		// @Router /api/v1/agreements [get]
		agmts.GET("", controller.GetAllAgreements)
	}
}
