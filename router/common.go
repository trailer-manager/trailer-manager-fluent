package router

import (
	"github.com/labstack/echo/v4"
	"github.com/trailer-manager/trailer-manager-fluent/controller"
)

func InitRouter(e *echo.Echo) {
	g := e.Group("/api/v1/trm")
	g.POST("/test/echo", controller.EchoTest)
	AddLogRouter(e)
}
