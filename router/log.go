package router

import (
	"github.com/labstack/echo/v4"
	"github.com/trailer-manager/trailer-manager-fluent/controller"
)

func AddLogRouter(e *echo.Echo, m ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/trm/log", m...)
	g.POST("/gps", controller.PostGpsLogData)
}
