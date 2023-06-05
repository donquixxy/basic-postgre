package handler

import (
	"postgre-basic/internal/api/request/userrequest"
	"postgre-basic/internal/exception"
	"postgre-basic/internal/responses"
	"postgre-basic/internal/usecases"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Create(c echo.Context) error
}

type UserHandlerImpl struct {
	Usecases usecases.UserServices
}

func NewUserHandler(
	Usecases usecases.UserServices,
) UserHandler {
	return &UserHandlerImpl{
		Usecases: Usecases,
	}
}

func (this *UserHandlerImpl) Create(c echo.Context) error {
	var respons responses.Response

	body, er := userrequest.ReadCreateRequest(c)
	er = this.Usecases.CreateUser(body)

	if er != nil {
		switch er.(type) {
		case *exception.BadRequestError:
			{
				respons.Error = "Bad Request"
				respons.ErrMsg = []string{er.Error()}
				respons.StatusCode = 400
				return c.JSON(respons.StatusCode, respons)
			}
		case *exception.DuplicateEntryError:
			{
				respons.Error = "Duplicate Entry"
				respons.ErrMsg = []string{er.Error()}
				respons.StatusCode = 409
				return c.JSON(respons.StatusCode, respons)
			}
		default:
			{
				respons.StatusCode = 503
				respons.Error = "Service error"
				respons.ErrMsg = []string{er.Error()}

				return c.JSON(respons.StatusCode, respons)
			}
		}
	}
	respons.StatusCode = 200
	respons.Data = body
	return c.JSON(respons.StatusCode, respons)
}
