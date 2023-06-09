package validator

import (
	"postgre-basic/internal/exception"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func ValidateStruct(value interface{}) ([]string, error) {
	validate := validator.New()
	var errString string
	var listErr []string
	err := validate.Struct(value)
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		for _, item := range err.(validator.ValidationErrors) {
			errString = item.Translate(trans)
			listErr = append(listErr, errString)
		}
		return listErr, &exception.BadRequestError{
			Message: "Invalid Input value",
		}
	}

	return nil, nil
}
