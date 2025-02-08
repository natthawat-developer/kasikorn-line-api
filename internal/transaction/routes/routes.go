package routes

import (
	"github.com/gofiber/fiber/v2"
	v1 "kasikorn-line-api/internal/transaction/handlers"
	"kasikorn-line-api/internal/transaction/services"
)

func RegisterRoutes(app *fiber.App, transactionService services.TransactionService) {

	v1TransactionHandler := v1.NewTransactionHandler(transactionService)
	v1TransactionRoutes := app.Group("/v1/transaction")
	v1TransactionRoutes.Get("/:transaction_id", v1TransactionHandler.GetTransactionDetail)
	v1TransactionRoutes.Get("/user/:user_id", v1TransactionHandler.GetTransactionByUserID)
}