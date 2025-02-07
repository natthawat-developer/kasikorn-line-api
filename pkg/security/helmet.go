package security

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)


func SetupHelmet(app *fiber.App) {
	app.Use(helmet.New())
}
