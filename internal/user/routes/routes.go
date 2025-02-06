package routes

import (
	"kasikorn-line-api/internal/user/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	// Register user-related routes
	userRoutes(app, userHandler)
}

func userRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	user := app.Group("/user") // Grouping routes for user

	// Route for getting user by user_id
	user.Get("/:user_id", userHandler.GetUser)
}
