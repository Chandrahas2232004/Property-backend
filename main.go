package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"property-backend/config"
	"property-backend/controllers"
	_ "property-backend/docs"
	"property-backend/repositories"
	"property-backend/routes"
	"property-backend/services"
)

func main() {

	/// ✅ INIT FILE LOGGING FIRST
	config.InitFileLogger()

	log.Println("🚀 Application starting...")

	// Initialize JWT (check if secret is configured)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Println("⚠️  JWT_SECRET not set - authentication tokens will fail")
	} else {
		log.Println("✅ JWT_SECRET configured")
	}

	// Initialize SendGrid for email functionality
	config.InitSendGrid()

	// Connect database
	config.ConnectDatabase()

	// Initialize S3 uploader
	config.InitS3()

	// setup gin
	r := gin.Default()
	db := config.DB

	// Setup CORS for integration purposes
	config.SetupCORS(r)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API versioning
	apiV1 := r.Group("/api/v1") // parent RouterGroup

	// Construct repositories (db may be nil)
	authRepo := repositories.NewAuthRepository(db)
	propertyRepo := repositories.NewPropertyRepository(db)
	agreementRepo := repositories.NewAgreementRepository(db)
	assetRepo := repositories.NewAssetRepository(db)
	contractRepo := repositories.NewContractRepository(db)
	formRepo := repositories.NewFormRepository(db)

	// Construct services
	authSvc := services.NewAuthService(authRepo)
	propertySvc := services.NewPropertyService(propertyRepo)
	agreementSvc := services.NewAgreementService(agreementRepo)
	assetSvc := services.NewAssetService(assetRepo)
	contractSvc := services.NewContractService(contractRepo)
	formSvc := services.NewFormService(formRepo)
	dashboardSvc := services.NewDashboardService(propertyRepo, assetRepo, contractRepo, agreementRepo)

	// Instantiate controllers with services
	authController := controllers.NewAuthController(authSvc)
	propertyController := controllers.NewPropertyController(propertySvc)
	agreementController := controllers.NewAgreementController(agreementSvc)
	assetController := controllers.NewAssetController(assetSvc)
	contractController := controllers.NewContractController(contractSvc)
	formController := controllers.NewFormController(formSvc)
	dashboardController := controllers.NewDashboardController(dashboardSvc)

	// Register routes under /api/v1
	routes.RegisterRoutes(
		apiV1,
		authController,
		propertyController,
		agreementController,
		assetController,
		contractController,
		formController,
		dashboardController,
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	log.Println("🚀 Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
