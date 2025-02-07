package health

import (
	"github.com/gofiber/fiber/v2"
	"kasikorn-line-api/pkg/database"
	coreError "kasikorn-line-api/pkg/error"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func RegisterRoutes(app *fiber.App) {

	app.Get("/health", HealthCheck)
}

func HealthCheck(c *fiber.Ctx) error {

	db, err := database.DB.DB()
	if err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusServiceUnavailable, "Database connection failed (unable to get DB object)")
	}

	if err := db.Ping(); err != nil {
		return coreError.HandleErrorResponse(c, fiber.StatusServiceUnavailable, "Database connection failed (Ping failed)")
	}

	response := HealthResponse{
		Status:  "success",
		Message: "Health check passed",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
