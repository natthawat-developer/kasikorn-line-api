package routes

import (
	"github.com/gofiber/fiber/v2"
	v1 "kasikorn-line-api/internal/account/handlers"
	"kasikorn-line-api/internal/account/services"
)

func RegisterRoutes(app *fiber.App, accountService services.AccountService) {

	v1AccountHandler := v1.NewAccountHandler(accountService)
	v1AccountRoutes := app.Group("/v1/account")
	v1AccountRoutes.Get("/:account_id", v1AccountHandler.GetAccountDetail)
	v1AccountRoutes.Get("/user/:user_id", v1AccountHandler.GetAccountByUserID)
	v1AccountRoutes.Get("/user/:user_id/main", v1AccountHandler.GetMainAccountByUserID)
}