package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kasikorn-line-api/internal/banner/models"
	"kasikorn-line-api/internal/banner/services"
	coreError "kasikorn-line-api/pkg/error"
	coreValidator "kasikorn-line-api/pkg/validator"
)

type BannerHandler struct {
	service services.BannerService
}

func NewBannerHandler(service services.BannerService) *BannerHandler {
	return &BannerHandler{service: service}
}

func (h *BannerHandler) GetBanner(c *fiber.Ctx) error {
	var req models.BannerRequest
	// Parse request parameters
	if err := c.ParamsParser(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, coreError.ErrInvalidParams)
	}

	// Validate the request
	if err := coreValidator.Validate(&req); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Get banner details from service
	banners, err := h.service.GetBannerDetails(req)
	if err != nil {
		// Check for custom error response
		if errorResponse, ok := err.(*coreError.ErrorResponse); ok {
			return c.Status(errorResponse.Code).JSON(errorResponse)
		}
		// Default error handling
		return coreError.HandleErrorResponse(c, fiber.StatusInternalServerError, coreError.ErrInternalServerError)
	}

	// Return the banners if no error
	return c.Status(fiber.StatusOK).JSON(banners)
}
