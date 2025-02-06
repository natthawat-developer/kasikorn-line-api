package middleware

import "github.com/gofiber/fiber/v2"

func ApplyMiddleware(app *fiber.App) {
	app.Use(logger.New())
	app.Use(recover.New())
}
