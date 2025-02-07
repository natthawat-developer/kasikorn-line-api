package routes

import (
	"kasikorn-line-api/internal/banner/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupBannerRoutes(app *fiber.App, bannerHandler *handlers.BannerHandler) {
	// Register banner-related routes
	bannerRoutes(app, bannerHandler)
}

func bannerRoutes(app *fiber.App, bannerHandler *handlers.BannerHandler) {
	banner := app.Group("/banner") // Grouping routes for banner

	// Route for getting banners by user_id
	banner.Get("/:user_id", bannerHandler.GetBanner)
}
