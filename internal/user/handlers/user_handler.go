package handlers

import (
	"kasikorn-line-api/internal/user/services"
	"kasikorn-line-api/internal/user/models"
	errpkg "kasikorn-line-api/pkg/error"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	// Use UserRequest from models
	var req models.UserRequest
	if err := c.ParamsParser(&req); err != nil { // ใช้ ParamsParser ในการดึงข้อมูลจาก URL params
		return c.Status(fiber.StatusBadRequest).JSON(&errpkg.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request parameters",
		})
	}

	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errpkg.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Get user from service
	user, err := h.service.GetUserDetails(req)
	if err != nil {
		// Check if the error is of type *errpkg.ErrorResponse
		// if errorResponse, ok := err.(*errpkg.ErrorResponse); ok {
		// 	return c.Status(errorResponse.Code).JSON(errorResponse)
		// }
		// Default error handling (for unexpected errors)
		defaultError := &errpkg.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "Something went wrong",
		}
		return c.Status(defaultError.Code).JSON(defaultError)
	}

	// Return the user details if no error
	return c.JSON(user)
}
