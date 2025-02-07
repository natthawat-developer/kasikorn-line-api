package main

import (
	"github.com/gofiber/fiber/v2"

	userRepos "kasikorn-line-api/internal/user/repositories"
	userRoutes "kasikorn-line-api/internal/user/routes"
	userServices "kasikorn-line-api/internal/user/services"

	bannerHandlers "kasikorn-line-api/internal/banner/handlers"
	bannerRepos "kasikorn-line-api/internal/banner/repositories"
	bannerRoutes "kasikorn-line-api/internal/banner/routes"
	bannerServices "kasikorn-line-api/internal/banner/services"

	"kasikorn-line-api/internal/config"
	"kasikorn-line-api/pkg/database"
	logger "kasikorn-line-api/pkg/log"
	"kasikorn-line-api/pkg/security"
)

func main() {
	// Initialize the logger
	logger.Initialize()
	defer logger.Close()

	// Load config
	appConfig := config.LoadConfig()

	// Initialize Fiber app
	app := fiber.New()

	// Set up CORS
	security.CorsSetup(app, security.Config{
		AllowOrigins: appConfig.CORS.AllowOrigins,
		AllowMethods: appConfig.CORS.AllowMethods,
		AllowHeaders: appConfig.CORS.AllowHeaders,
	})

	// Connect to the database
	if err := database.Connect(database.Config{
		User:     appConfig.DB.User,
		Password: appConfig.DB.Password,
		Host:     appConfig.DB.Host,
		Port:     appConfig.DB.Port,
		Name:     appConfig.DB.Name,
		Logger:   logger.NewZapGormLogger(logger.Logger),
	}); err != nil {
		logger.Error("Failed to connect to the database")
	}

	userRepo := userRepos.NewUserRepository(database.DB)
	userService := userServices.NewUserService(userRepo)
	userRoutes.RegisterRoutes(app, userService)

	bannerRepo := bannerRepos.NewBannerRepository(database.DB)
	bannerService := bannerServices.NewBannerService(bannerRepo)
	bannerHandler := bannerHandlers.NewBannerHandler(bannerService)

	bannerRoutes.SetupBannerRoutes(app, bannerHandler)
	// Start the server
	if err := app.Listen(":" + appConfig.Port); err != nil {
		logger.Fatal("Failed to start server")
	}
}
