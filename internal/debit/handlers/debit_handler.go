package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kasikorn-line-api/internal/debit/models"
	"kasikorn-line-api/internal/debit/services"
	coreError "kasikorn-line-api/pkg/error"
	coreValidator "kasikorn-line-api/pkg/validator"
)

type DebitHandler struct {
	service services.DebitService
}

func NewDebitHandler(service services.DebitService) *DebitHandler {
	return &DebitHandler{service: service}
}


func (h *DebitHandler) GetDebitDetail(c *fiber.Ctx) error {
	var req models.GetDebitCardDetailsByUserIDRequest
	// Parse request parameters
	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	// Validate the request
	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Get debit details from service
	debit, err := h.service.GetDebitCardDetailsByUserID(req)
	if err != nil {
		// Check for custom error response
		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}
		// Default error handling
		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	// Return the debit details if no error
	return c.Status(fiber.StatusOK).JSON(debit)
}

