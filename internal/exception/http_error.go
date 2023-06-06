package exception

import (
	"postgre-basic/internal/responses"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func ErrorRouteHandler(err error, c echo.Context) {
	log.Error("Error wrong route handler", err.Error())
	res := responses.Response{
		StatusCode: 523,
		Message:    "Invalid Route",
		ErrMsg:     []string{"Wrong route"},
		Data:       nil,
	}
	c.JSON(500, res)
}
