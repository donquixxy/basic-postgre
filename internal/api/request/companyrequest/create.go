package companyrequest

import (
	"postgre-basic/internal/exception"

	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	Name  string `json:"name" form:"name" validate:"required,min=5,max=20"`
	Phone string `json:"phone" form:"phone" validate:"required,min=9,max=16,numeric"`
}

func ReadCreateRequest(c echo.Context) (*CreateRequest, error) {
	body := new(CreateRequest)

	bindErr := c.Bind(body)

	if bindErr != nil {
		return nil, &exception.BadRequestError{Message: bindErr.Error()}
	}

	return body, nil
}
