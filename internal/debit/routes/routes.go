package routes

import (
	"github.com/gofiber/fiber/v2"
	v1 "kasikorn-line-api/internal/debit/handlers"
	"kasikorn-line-api/internal/debit/services"
)

func RegisterRoutes(app *fiber.App, debitService services.DebitService) {

	v1DebitHandler := v1.NewDebitHandler(debitService)
	v1DebitRoutes := app.Group("/v1/debit")
	v1DebitRoutes.Get("/user/:user_id", v1DebitHandler.GetDebitDetail)
}