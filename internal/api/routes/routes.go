package routes

import (
	"postgre-basic/internal/handler"

	"github.com/labstack/echo/v4"
)

func UserRoutes(c *echo.Echo, handler handler.UserHandler) {
	group := c.Group("/api/v1")
	group.POST("/user/create", handler.Create)
}
