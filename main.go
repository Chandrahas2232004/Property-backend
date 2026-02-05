package main

import (
	"io"
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
	// setup logging to stdout and app.log
	logFile, err := os.OpenFile("docs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("could not open log file: %v", err)
	} else {
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
	}

	// Initialize DB via GORM and config
	config.ConnectDatabase()
	db := config.DB

	r := gin.Default()

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

	// Construct services
	authSvc := services.NewAuthService(authRepo)
	propertySvc := services.NewPropertyService(propertyRepo)
	agreementSvc := services.NewAgreementService(agreementRepo)
	assetSvc := services.NewAssetService(assetRepo)
	contractSvc := services.NewContractService(contractRepo)

	// Instantiate controllers with services
	authController := controllers.NewAuthController(authSvc)
	propertyController := controllers.NewPropertyController(propertySvc)
	agreementController := controllers.NewAgreementController(agreementSvc)
	assetController := controllers.NewAssetController(assetSvc)
	contractController := controllers.NewContractController(contractSvc)

	// Register routes under /api/v1
	routes.RegisterRoutes(
		apiV1,
		authController,
		propertyController,
		agreementController,
		assetController,
		contractController,
	)

	
	// Start server
	r.Run(":8080")
}
