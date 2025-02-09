package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func UUIDValidation(fl validator.FieldLevel) bool {

	regex := `^[0-9a-fA-F]{8}[0-9a-fA-F]{4}[0-9a-fA-F]{4}[0-9a-fA-F]{4}[0-9a-fA-F]{12}$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(fl.Field().String())
}

func Validate(data interface{}) error {
	validate := validator.New()

	validate.RegisterValidation("uuid", UUIDValidation)

	return validate.Struct(data)
}
