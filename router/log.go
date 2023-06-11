package router

import (
	"SiverPineValley/trailer-manager/controller"
	"github.com/labstack/echo/v4"
)

func AddLogRouter(e *echo.Echo, m ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/trm/log", m...)
	g.POST("/gps", controller.PostGpsLogData)
}