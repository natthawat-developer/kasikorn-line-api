package validator

import (
	"regexp"
	"github.com/go-playground/validator/v10"
)

// UUIDValidation function สำหรับการตรวจสอบ UUID
func UUIDValidation(fl validator.FieldLevel) bool {
	// Define a regex pattern for UUID format with or without dash
	regex := `^[0-9a-fA-F]{8}[0-9a-fA-F]{4}[0-9a-fA-F]{4}[0-9a-fA-F]{4}[0-9a-fA-F]{12}$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(fl.Field().String())
}

// Validate function สำหรับการ validate struct ใดๆ
func Validate(data interface{}) error {
	validate := validator.New()

	// Register custom validation
	validate.RegisterValidation("uuid", UUIDValidation)

	// Validate the request
	return validate.Struct(data)
}
