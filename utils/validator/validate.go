package validator

import (
	"postgre-basic/internal/exception"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(value interface{}) ([]string, error) {
	validate := validator.New()
	var errString string
	var listErr []string
	err := validate.Struct(value)

	if err != nil {
		for _, item := range err.(validator.ValidationErrors) {
			errString = item.Field() + " is " + item.Tag()
			listErr = append(listErr, errString)
		}
		return listErr, &exception.BadRequestError{
			Message: "Required field cant be empty",
		}
	}

	return nil, nil
}
