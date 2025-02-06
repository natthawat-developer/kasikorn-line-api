package error

import "github.com/gofiber/fiber/v2"

// Error messages
const (
	ErrInvalidParams      = "Invalid request parameters"
	ErrInternalServerError   = "Internal server error"
	ErrNotFound = "Record not found"
	ErrUnauthorizedAccess = "Unauthorized access"
)

// ErrorResponse represents the custom error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse returns a new ErrorResponse
func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

// Implement the error interface
func (e *ErrorResponse) Error() string {
	return e.Message
}


func HandleErrorResponse(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(&ErrorResponse{
		Code:    code,
		Message: message,
	})
}