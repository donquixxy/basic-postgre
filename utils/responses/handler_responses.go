package responses

import (
	"postgre-basic/internal/exception"
	"postgre-basic/internal/responses"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	badRequestStatusCode     = 400
	duplicateEntryStatusCode = 409
	notFoundStatusCode       = 404
	serviceErrorStatusCode   = 503
	succesStatusCode         = 200
)

func GetReturnData(er error, c echo.Context, data interface{}, errMsg []string) error {

	respons := &responses.Response{}

	if er == nil {
		respons.StatusCode = succesStatusCode
		respons.Data = data
		respons.Message = "Success"
		return c.JSON(respons.StatusCode, respons)
	}

	log.Error(`Error :`, er.Error())
	if er != nil {
		switch er.(type) {
		case *exception.BadRequestError:
			{
				respons.Message = "Bad Request"
				respons.ErrMsg = errMsg
				respons.ErrMsg = append(respons.ErrMsg, er.Error())
				respons.StatusCode = badRequestStatusCode
				return c.JSON(respons.StatusCode, respons)
			}
		case *exception.DuplicateEntryError:
			{
				respons.Message = "Duplicate Entry"
				respons.ErrMsg = errMsg
				respons.ErrMsg = append(respons.ErrMsg, er.Error())
				respons.StatusCode = duplicateEntryStatusCode
				return c.JSON(respons.StatusCode, respons)
			}
		case *exception.RecordNotFoundError:
			{
				respons.Message = "Record Not Found"
				respons.ErrMsg = errMsg
				respons.ErrMsg = append(respons.ErrMsg, er.Error())
				respons.StatusCode = notFoundStatusCode
				return c.JSON(respons.StatusCode, respons)
			}
		default:
			{
				respons.StatusCode = serviceErrorStatusCode
				respons.Message = "Service error"
				respons.ErrMsg = errMsg
				respons.ErrMsg = append(respons.ErrMsg, er.Error())
				return c.JSON(respons.StatusCode, respons)
			}
		}
	}

	return nil
}
