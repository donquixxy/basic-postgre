package userrequest

import (
	"fmt"
	"postgre-basic/internal/exception"

	"github.com/labstack/echo/v4"
)

type UpdateRequest struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}

func ReadUpdateRequest(c echo.Context) (*UpdateRequest, error) {
	payload := new(UpdateRequest)

	bindErr := c.Bind(payload)

	if bindErr != nil {
		fmt.Println("Error binding Updaterequest: ", bindErr.Error())
		return nil, &exception.BadRequestError{
			Message: bindErr.Error(),
		}
	}

	return payload, nil
}
