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
	// Parse request parameters
	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	// Validate the request
	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Get transaction details from service
	res, err := h.service.GetTransactionByUserID(req)
	if err != nil {
		// Check for custom error response
		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}
		// Default error handling
		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	// Return the transaction details if no error
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *TransactionHandler) GetTransactionDetail(c *fiber.Ctx) error {
	var req models.GetTransactionDetailRequest
	// Parse request parameters
	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	// Parse query parameters (query params)
	if err := c.QueryParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}
	
	// Validate the request
	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Get transaction details from service
	res, err := h.service.GetTransactionDetail(req)
	if err != nil {
		// Check for custom error response
		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}
		// Default error handling
		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	// Return the transaction details if no error
	return c.Status(fiber.StatusOK).JSON(res)
}
