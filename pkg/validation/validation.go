package validation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
)

var uni *ut.UniversalTranslator
var validate = validator.New()
var Translator ut.Translator

func init() {
	enTrans := en.New()
	uni = ut.New(enTrans, enTrans)
	Translator, _ = uni.GetTranslator("en")

	_ = validate.RegisterTranslation("required", Translator, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required field.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
}

func New() *validator.Validate {
	return validator.New()
}

// GetValidator returns instance of validator
func GetValidator() *validator.Validate {
	return validate
}
