package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCORS configures CORS for the Gin engine
// This allows cross-origin requests from specified origins for integration purposes
func SetupCORS(r *gin.Engine) {
	corsConfig := cors.Config{
		// AllowOrigins: []string{"http://localhost:3000", "http://localhost:5173"},  // for development
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-CSRF-Token",
			"X-Requested-With",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Type",
			"X-Total-Count",

			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           3600, // 1 hour
	}

	r.Use(cors.New(corsConfig))
}
