package router

import (
	"SiverPineValley/trailer-manager/controller"
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	g := e.Group("/api/v1/trm")
	g.POST("/test/echo", controller.EchoTest)
	AddLogRouter(e)
}