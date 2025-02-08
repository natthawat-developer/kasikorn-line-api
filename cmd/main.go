package main

import (
	"time"

	"github.com/gofiber/fiber/v2"

	userRepos "kasikorn-line-api/internal/user/repositories"
	userRoutes "kasikorn-line-api/internal/user/routes"
	userServices "kasikorn-line-api/internal/user/services"

	bannerRepos "kasikorn-line-api/internal/banner/repositories"
	bannerRoutes "kasikorn-line-api/internal/banner/routes"
	bannerServices "kasikorn-line-api/internal/banner/services"

	accountRepos "kasikorn-line-api/internal/account/repositories"
	accountRoutes "kasikorn-line-api/internal/account/routes"
	accountServices "kasikorn-line-api/internal/account/services"

	debitRepos "kasikorn-line-api/internal/debit/repositories"
	debitRoutes "kasikorn-line-api/internal/debit/routes"
	debitServices "kasikorn-line-api/internal/debit/services"

	transactionRepos "kasikorn-line-api/internal/transaction/repositories"
	transactionRoutes "kasikorn-line-api/internal/transaction/routes"
	transactionServices "kasikorn-line-api/internal/transaction/services"

	"kasikorn-line-api/internal/config"
	"kasikorn-line-api/pkg/database"
	"kasikorn-line-api/pkg/health"
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
	security.CorsSetup(app, security.CorsConfig{
		AllowOrigins: appConfig.CORS.AllowOrigins,
		AllowMethods: appConfig.CORS.AllowMethods,
		AllowHeaders: appConfig.CORS.AllowHeaders,
	})

	// Set up Rate Limiting
	security.SetupRateLimiter(app, security.RateLimiterConfig{
		Max:        appConfig.RateLimiter.MaxRequests,
		Expiration: time.Duration(appConfig.RateLimiter.Expiration) * time.Second,
	})

	security.SetupHelmet(app)

	// Connect to the database
	if err := database.Connect(database.DatabaseConfig{
		User:     appConfig.DB.User,
		Password: appConfig.DB.Password,
		Host:     appConfig.DB.Host,
		Port:     appConfig.DB.Port,
		Name:     appConfig.DB.Name,
		Logger:   logger.NewZapGormLogger(logger.Logger),
	}); err != nil {
		logger.Error("Failed to connect to the database")
	}

	// Register health check route
	health.RegisterRoutes(app)

	userRepo := userRepos.NewUserRepository(database.DB)
	userService := userServices.NewUserService(userRepo)
	userRoutes.RegisterRoutes(app, userService)

	bannerRepo := bannerRepos.NewBannerRepository(database.DB)
	bannerService := bannerServices.NewBannerService(bannerRepo)
	bannerRoutes.RegisterRoutes(app, bannerService)

	accountRepo := accountRepos.NewAccountRepository(database.DB)
	accountService := accountServices.NewAccountService(accountRepo)
	accountRoutes.RegisterRoutes(app, accountService)

	debitRepo := debitRepos.NewDebitRepository(database.DB)
	debitService := debitServices.NewDebitService(debitRepo)
	debitRoutes.RegisterRoutes(app, debitService)

	transactionRepo := transactionRepos.NewTransactionRepository(database.DB)
	transactionService := transactionServices.NewTransactionService(transactionRepo)
	transactionRoutes.RegisterRoutes(app, transactionService)

	// Start the server
	if err := app.Listen(":" + appConfig.Port); err != nil {
		logger.Fatal("Failed to start server")
	}
}
