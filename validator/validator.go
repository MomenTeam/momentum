package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validator variable
var Validator *validator.Validate

// Setup validator
func Setup() {
	Validator = validator.New()
	Validator.RegisterValidation("status", ValidateStatus)
}

// ValidateStatus function
func ValidateStatus(fl validator.FieldLevel) bool {
	return fl.Field().String() == "PENDING" || fl.Field().String() == "PROCESSING" || fl.Field().String() == "DELIVERED"
}
