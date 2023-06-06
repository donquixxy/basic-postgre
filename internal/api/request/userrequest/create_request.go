package userrequest

import (
	"fmt"
	"postgre-basic/internal/exception"

	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
	Age  int    `json:"age" form:"age" validate:"required"`
}

func ReadCreateRequest(c echo.Context) (*CreateRequest, error) {
	payload := new(CreateRequest)

	bindErr := c.Bind(payload)

	if bindErr != nil {
		fmt.Println("Error binding CreateRequest: ", bindErr.Error())
		return nil, &exception.BadRequestError{
			Message: bindErr.Error(),
		}
	}

	return payload, nil

}
