package handler

import (
	"postgre-basic/internal/api/request/companyrequest"
	"postgre-basic/internal/usecases"
	"postgre-basic/utils/responses"
	"postgre-basic/utils/validator"

	"github.com/labstack/echo/v4"
)

type CompanyHandler interface {
	Create(c echo.Context) error
}

type CompanyHandlerImpl struct {
	Usecase usecases.CompanyServices
}

func NewCompanyHandler(
	Usecase usecases.CompanyServices,
) CompanyHandler {
	return &CompanyHandlerImpl{
		Usecase: Usecase,
	}
}

func (this *CompanyHandlerImpl) Create(c echo.Context) error {
	body, err := companyrequest.ReadCreateRequest(c)

	if err != nil {
		return responses.GetReturnData(err, c, nil, []string{})
	}

	validated, err := validator.ValidateStruct(body)

	if err != nil {
		return responses.GetReturnData(err, c, nil, validated)
	}

	payload, err := this.Usecase.Create(body)

	if err != nil {
		return responses.GetReturnData(err, c, nil, nil)
	}

	return responses.GetReturnData(nil, c, payload, nil)
}
