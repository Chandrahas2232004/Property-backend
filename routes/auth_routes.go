package routes

import (
	"property-backend/controllers"

	"github.com/gin-gonic/gin"
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

		// ForgotPassword
		// @Summary Request a password reset
		// @Tags Auth
		// @Accept json
		// @Produce json
		// @Router /api/v1/auth/forgot-password [post]
		auth.POST("/forgot-password", controller.ForgotPassword)

		// ResetPassword
		// @Summary Reset password with token
		// @Tags Auth
		// @Accept json
		// @Produce json
		// @Router /api/v1/auth/reset-password [post]
		auth.POST("/reset-password", controller.ResetPassword)
	}
}
