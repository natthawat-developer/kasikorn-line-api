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
	logger.Initialize()  // เริ่มต้น logger
	defer logger.Close() // ปิด logger เมื่อแอปปิดตัวลง

	// Load config
	appConfig := config.LoadConfig()

	// Initialize Fiber app
	app := fiber.New()

	// Connect to the database
	if err := database.Connect(
		appConfig.DB.User,
		appConfig.DB.Password,
		appConfig.DB.Host,
		appConfig.DB.Port,
		appConfig.DB.Name,
	); err != nil {
		// ใช้ logger ที่เราสร้างใน pkg/log
		logger.Error("Failed to connect to the database") // ใช้ logger แทนการเรียก zap โดยตรง
	}

	userRepo := userrepos.NewUserRepository(database.DB)
	userService := userservices.NewUserService(userRepo)
	userHandler := userhandlers.NewUserHandler(userService)

	userroutes.SetupUserRoutes(app, userHandler)

	// Start the server
	if err := app.Listen(":" + appConfig.Port); err != nil {
		// ใช้ logger ที่เราสร้างใน pkg/log
		logger.Fatal("Failed to start server") // ใช้ logger แทนการเรียก zap โดยตรง
	}
}
