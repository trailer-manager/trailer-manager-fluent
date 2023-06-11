package controller

import (
	"SiverPineValley/trailer-manager/logger"
	"SiverPineValley/trailer-manager/service"
	"SiverPineValley/trailer-manager/utility"
	"net/http"

	"github.com/labstack/echo/v4"
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
