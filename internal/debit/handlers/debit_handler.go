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

func (h *DebitHandler) GetDebitCardsByUserID(c *fiber.Ctx) error {
	var req models.GetDebitCardsByUserIDRequest

	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.service.GetDebitCardsByUserID(req)
	if err != nil {

		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}

		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *DebitHandler) GetDebitDetail(c *fiber.Ctx) error {
	var req models.GetDebitCardDetailsByCardIDRequest

	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	if err := c.QueryParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.service.GetDebitCardDetailsByCardID(req)
	if err != nil {

		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}

		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
