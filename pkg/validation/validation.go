package validation

import "github.com/go-playground/validator"

var validate = validator.New()

func New() *validator.Validate {
	return validator.New()
}

// GetValidator returns instance of validator
func GetValidator() *validator.Validate {
	return validate
}
