package routes

import (
	"postgre-basic/internal/handler"

	"github.com/labstack/echo/v4"
)

func UserRoutes(c *echo.Echo, handler handler.UserHandler) {
	group := c.Group("/api/v1")
	group.POST("/user/create", handler.Create)
	group.GET("/user", handler.FindAll)
	group.PUT("/user/:id/update", handler.Update)
	group.GET("/user/:id/details", handler.FindByID)
	group.DELETE("/user/:id/delete", handler.Delete)
}

func CompanyRoute(c *echo.Echo, handler handler.CompanyHandler) {
	group := c.Group("/api/v1")
	group.POST("/company/create", handler.Create)
}
