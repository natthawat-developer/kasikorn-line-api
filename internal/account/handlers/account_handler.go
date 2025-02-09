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

	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	account, err := h.service.GetAccountDetail(req)
	if err != nil {

		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}

		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(account)
}

func (h *AccountHandler) GetAccountByUserID(c *fiber.Ctx) error {
	var req models.GetAccountByUserIDRequest

	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	account, err := h.service.GetAccountByUserID(req)
	if err != nil {

		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}

		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(account)
}

func (h *AccountHandler) GetMainAccountByUserID(c *fiber.Ctx) error {
	var req models.GetMainAccountByUserIDRequest

	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	account, err := h.service.GetMainAccountByUserID(req)
	if err != nil {

		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}

		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(account)
}
