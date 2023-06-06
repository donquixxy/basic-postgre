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
	FindAll(c echo.Context) error
	Update(c echo.Context) error
	FindByID(c echo.Context) error
	Delete(c echo.Context) error
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

func (this *UserHandlerImpl) FindAll(c echo.Context) error {
	listUsers, er := this.Usecases.FindAllUsers()

	if er != nil {
		return responses.GetReturnData(er, c, nil, []string{})
	}

	return responses.GetReturnData(er, c, listUsers, []string{})
}

func (this *UserHandlerImpl) Update(c echo.Context) error {
	id := c.Param("id")
	body, err := userrequest.ReadUpdateRequest(c)

	if err != nil {
		return responses.GetReturnData(err, c, nil, []string{})
	}

	payload, err := this.Usecases.UpdateUsers(body, id)

	if err != nil {
		return responses.GetReturnData(err, c, nil, []string{})
	}

	return responses.GetReturnData(err, c, payload, []string{})
}

func (this *UserHandlerImpl) FindByID(c echo.Context) error {
	id := c.Param("id")
	payload, err := this.Usecases.FindByID(id)

	if err != nil {
		return responses.GetReturnData(err, c, nil, []string{})
	}

	return responses.GetReturnData(err, c, payload, []string{})
}

func (this *UserHandlerImpl) Delete(c echo.Context) error {
	id := c.Param("id")
	err := this.Usecases.Delete(id)

	if err != nil {
		return responses.GetReturnData(err, c, nil, []string{})
	}

	return responses.GetReturnData(err, c, "Success Delete data !", []string{})
}
