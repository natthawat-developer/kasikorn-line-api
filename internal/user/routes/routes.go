package routes

import (
	"github.com/gofiber/fiber/v2"
	v1 "kasikorn-line-api/internal/user/handlers"
	"kasikorn-line-api/internal/user/services"
)

func RegisterRoutes(app *fiber.App, userService services.UserService) {

	v1UserHandler := v1.NewUserHandler(userService)
	v1UserRoutes := app.Group("/v1/user")
	v1UserRoutes.Get("/:user_id", v1UserHandler.GetUserDetails)

}
