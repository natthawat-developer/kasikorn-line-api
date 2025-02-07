package security

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type CorsConfig struct {
	AllowOrigins string
	AllowMethods string
	AllowHeaders string
}

func CorsSetup(app *fiber.App, corsConfig CorsConfig) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: corsConfig.AllowOrigins,
		AllowMethods: corsConfig.AllowMethods,
		AllowHeaders: corsConfig.AllowHeaders,
	}))
}
