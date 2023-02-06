package api

import (
	"SiverPineValley/trailer-manager/common"
	"github.com/labstack/echo/v4"
)

func NotFoundHandler(c echo.Context) error {
	err := common.GetWebError(c, common.WebErrNotFound)
	c.Error(err)
	return err
}