package routes

import (
	"github.com/gin-gonic/gin"
	"property-backend/controllers"
)

// ContractRoutes registers contract endpoints under /contracts
func ContractRoutes(rg *gin.RouterGroup, controller *controllers.ContractController) {
	contracts := rg.Group("/contracts")
	{
		// Add contract
		// @Summary Create a contract
		// @Tags Contracts
		// @Accept json
		// @Produce json
		// @Router /api/v1/contracts [post]
		contracts.POST("", controller.AddContract)

		// Get all contracts
		// @Summary Get all contracts
		// @Tags Contracts
		// @Produce json
		// @Router /api/v1/contracts [get]
		contracts.GET("", controller.GetAllContracts)

		// Lease contracts
		// @Summary Get all lease contracts
		// @Tags Contracts
		// @Produce json
		// @Router /api/v1/contracts/lease [get]
		contracts.GET("/lease", controller.GetLeaseContracts)

		// AMC contracts
		// @Summary Get all AMC contracts
		// @Tags Contracts
		// @Produce json
		// @Router /api/v1/contracts/amc [get]
		contracts.GET("/amc", controller.GetAMCContracts)
	}
}
