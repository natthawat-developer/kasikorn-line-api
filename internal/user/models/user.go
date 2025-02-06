package models

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// UserRequest struct พร้อมกับ tag validation
type UserRequest struct {
	UserID string `params:"user_id" validate:"required,uuid"`
}
// UserResponse struct
type UserResponse struct {
    UserID   string `json:"user_id"`
    Name     string `json:"name"`
}

// Validate function สำหรับ UserRequest
func (req *UserRequest) Validate() error {
	validate := validator.New()

	// Register custom validation for UUID
	validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		// Define a regex pattern for UUID format with or without dash
		regex := `^[0-9a-fA-F]{8}[0-9a-fA-F]{4}[0-9a-fA-F]{4}[0-9a-fA-F]{4}[0-9a-fA-F]{12}$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
		re := regexp.MustCompile(regex)
		return re.MatchString(fl.Field().String())
	})

	return validate.Struct(req)
}