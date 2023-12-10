package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/trailer-manager/trailer-manager-common/logger"
	"github.com/trailer-manager/trailer-manager-common/utility"
	"github.com/trailer-manager/trailer-manager-fluent/service"
	"net/http"
)

// PostGpsLogData는 gps, 단말기, 와이파이 스캐닝 정보를 받아서 DB에 적재하는 함수
func PostGpsLogData(c echo.Context) error {
	tErr := service.PostGpsLog(c)
	if tErr.IsError() {
		logger.ErrorContext(utility.GetContextFromEchoContext(c), tErr.Error())
		responseJson(c, http.StatusInternalServerError, tErr)
		return tErr
	}

	responseJson(c, http.StatusOK, nil)
	return nil
}
