package routes

import (
	"github.com/gofiber/fiber/v2"
	v1 "kasikorn-line-api/internal/banner/handlers"
	"kasikorn-line-api/internal/banner/services"
)

func RegisterRoutes(app *fiber.App, bannerService services.BannerService) {

	v1BannerHandler := v1.NewBannerHandler(bannerService)
	v1BannerRoutes := app.Group("/v1/banner")
	v1BannerRoutes.Get("/:user_id", v1BannerHandler.GetBanner)

}
