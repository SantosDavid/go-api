package presenter

import (
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func NewValidationError(validationError error) ErrorPresenter {
	errors := make(map[string]interface{})

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	for _, e := range validationError.(validator.ValidationErrors) {
		errors[strings.ToLower(e.Field())] = e.Translate(trans)
	}

	return ErrorPresenter{Msg: "validation errors", Errors: errors}
}
