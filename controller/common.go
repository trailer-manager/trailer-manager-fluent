package controller

import (
	model "github.com/trailer-manager/trailer-manager-fluent/model/api"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

func responseJson(c echo.Context, statusCode int, resBody interface{}) {
	response := model.Response{
		Result: model.Result{
			StatusCode:    statusCode,
			TransactionId: c.Request().Header.Get("transactionId"),
		},
	}

	if resBody != nil {
		copier.Copy(resBody, response)
		c.JSON(statusCode, resBody)
		return
	}

	c.JSON(statusCode, response)
	return
}
