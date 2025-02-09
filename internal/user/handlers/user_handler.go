package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kasikorn-line-api/internal/user/models"
	"kasikorn-line-api/internal/user/services"
	coreError "kasikorn-line-api/pkg/error"
	coreValidator "kasikorn-line-api/pkg/validator"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserDetails(c *fiber.Ctx) error {
	var req models.UserRequest

	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := h.service.GetUserDetails(req)
	if err != nil {

		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}

		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
