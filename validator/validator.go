package validator

import (
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = validator.New()
	Validator.RegisterValidation("department", ValidateDepartment)
}

func ValidateDepartment(fl validator.FieldLevel) bool {
	return fl.Field().String() == "Design" || fl.Field().String() == "Marketing" || fl.Field().String() == "Development"
}
