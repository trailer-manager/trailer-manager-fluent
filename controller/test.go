package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// EchoTest is testing function.
func EchoTest(c echo.Context) error {
	var resBody any
	err := c.Bind(&resBody); if err != nil {
		responseJson(c, http.StatusBadRequest, nil)
		return err
	}

	responseJson(c, http.StatusOK, resBody)
	return nil
}