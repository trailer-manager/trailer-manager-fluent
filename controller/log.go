package controller

import (
	"SiverPineValley/trailer-manager/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

// PostGpsLogData는 gps, 단말기, 와이파이 스캐닝 정보를 받아서 DB에 적재하는 함수
func PostGpsLogData(c echo.Context) error {
	tErr := service.PostGpsLog(c)
	if tErr.IsError() {
		responseJson(c, http.StatusInternalServerError, tErr.Error())
		return tErr
	}

	return nil
}
