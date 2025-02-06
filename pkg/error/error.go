package error

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