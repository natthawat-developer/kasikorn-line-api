package main

import (
	"github.com/gofiber/fiber/v2"

	userhandlers "kasikorn-line-api/internal/user/handlers"
	userrepos "kasikorn-line-api/internal/user/repositories"
	userroutes "kasikorn-line-api/internal/user/routes"
	userservices "kasikorn-line-api/internal/user/services"

	"kasikorn-line-api/internal/config"
	"kasikorn-line-api/pkg/database"
	logger "kasikorn-line-api/pkg/log"
)

func main() {
	// Initialize the logger
	logger.Initialize()  
	defer logger.Close() 

	// Load config
	appConfig := config.LoadConfig()

	// Initialize Fiber app
	app := fiber.New()

	dbConfig := database.Config{
		User:     appConfig.DB.User,
		Password: appConfig.DB.Password,
		Host:     appConfig.DB.Host,
		Port:     appConfig.DB.Port,
		Name:     appConfig.DB.Name,
		Logger:   logger.NewZapGormLogger(logger.Logger),
	}

	// Connect to the database
	if err := database.Connect(dbConfig); err != nil {
		logger.Error("Failed to connect to the database")
	}

	userRepo := userrepos.NewUserRepository(database.DB)
	userService := userservices.NewUserService(userRepo)
	userHandler := userhandlers.NewUserHandler(userService)

	userroutes.SetupUserRoutes(app, userHandler)

	// Start the server
	if err := app.Listen(":" + appConfig.Port); err != nil {
		logger.Fatal("Failed to start server") 
	}
}
