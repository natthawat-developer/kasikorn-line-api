package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kasikorn-line-api/internal/transaction/models"
	"kasikorn-line-api/internal/transaction/services"
	coreError "kasikorn-line-api/pkg/error"
	coreValidator "kasikorn-line-api/pkg/validator"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) GetTransactionByUserID(c *fiber.Ctx) error {
	var req models.GetTransactionByUserIDRequest

	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.service.GetTransactionByUserID(req)
	if err != nil {

		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}

		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *TransactionHandler) GetTransactionDetail(c *fiber.Ctx) error {
	var req models.GetTransactionDetailRequest

	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	if err := c.QueryParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.service.GetTransactionDetail(req)
	if err != nil {

		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}

		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
