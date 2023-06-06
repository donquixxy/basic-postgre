package handler

import (
	"postgre-basic/internal/api/request/userrequest"
	"postgre-basic/internal/usecases"
	"postgre-basic/utils/responses"
	"postgre-basic/utils/validator"

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

	body, er := userrequest.ReadCreateRequest(c)

	if er != nil {
		return responses.GetReturnData(er, c, nil, []string{})
	}

	errMsg, er := validator.ValidateStruct(body)

	if er != nil {
		return responses.GetReturnData(er, c, nil, errMsg)
	}

	data, er := this.Usecases.CreateUser(body)

	if er != nil {
		return responses.GetReturnData(er, c, nil, []string{})
	}

	return responses.GetReturnData(er, c, data, []string{})
}
