package routes

import (
	"github.com/gin-gonic/gin"
	"property-backend/controllers"
)

// AuthRoutes registers authentication routes under /auth
func AuthRoutes(rg *gin.RouterGroup, controller *controllers.AuthController) {
	auth := rg.Group("/auth")
	{
		// SignUp
		// @Summary Sign up a new user
		// @Tags Auth
		// @Accept json
		// @Produce json
		// @Router /api/v1/auth/signup [post]
		auth.POST("/signup", controller.SignUp)

		// SignIn
		// @Summary Sign in a user
		// @Tags Auth
		// @Accept json
		// @Produce json
		// @Router /api/v1/auth/signin [post]
		auth.POST("/signin", controller.SignIn)
	}
}
