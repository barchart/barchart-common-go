package validation

import "github.com/go-playground/validator"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// GetValidator returns instance of validator
func GetValidator() *validator.Validate {
	return validate
}
