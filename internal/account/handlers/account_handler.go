package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kasikorn-line-api/internal/account/models"
	"kasikorn-line-api/internal/account/services"
	coreError "kasikorn-line-api/pkg/error"
	coreValidator "kasikorn-line-api/pkg/validator"
)

type AccountHandler struct {
	service services.AccountService
}

func NewAccountHandler(service services.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}


func (h *AccountHandler) GetAccountDetail(c *fiber.Ctx) error {
	var req models.GetAccountDetailRequest
	// Parse request parameters
	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	// Validate the request
	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Get account details from service
	account, err := h.service.GetAccountDetail(req)
	if err != nil {
		// Check for custom error response
		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}
		// Default error handling
		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	// Return the account details if no error
	return c.Status(fiber.StatusOK).JSON(account)
}

func (h *AccountHandler) GetAccountByUserID(c *fiber.Ctx) error {
	var req models.GetAccountByUserIDRequest
	// Parse request parameters
	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	// Validate the request
	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Get account details from service
	account, err := h.service.GetAccountByUserID(req)
	if err != nil {
		// Check for custom error response
		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}
		// Default error handling
		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	// Return the account details if no error
	return c.Status(fiber.StatusOK).JSON(account)
}