package routes

import (
	"github.com/gin-gonic/gin"
	"property-backend/controllers"
)

func FormRoutes(rg *gin.RouterGroup, controller *controllers.FormController) {
	forms := rg.Group("/form")
	{
		forms.GET("/country", controller.GetCountry)
		forms.GET("/state/:id", controller.GetState)
		forms.GET("/district/:id", controller.GetDistrict)
		forms.GET("/taluk/:id", controller.GetTaluk)
	}
}