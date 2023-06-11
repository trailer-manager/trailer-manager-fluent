package service

import (
	"SiverPineValley/trailer-manager/common"
	db "SiverPineValley/trailer-manager/db/rdb"
	"SiverPineValley/trailer-manager/logger"
	model "SiverPineValley/trailer-manager/model/api"

	"SiverPineValley/trailer-manager/utility"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func PostGpsLog(c echo.Context) (trmErr *common.Error) {
	ctx := utility.GetContextFromEchoContext(c)

	gpsLog := model.GpsLog{}
	if err := c.Bind(&gpsLog); err != nil {
		logger.ErrorContext(ctx, err.Error())
		return common.NewServerError(c, common.ServerErrBadRequest)
	}

	store := db.NewStore(db.RDB)
	_, err := store.CreateGpsLog(ctx, db.CreateGpsLogParams{
		Sid:     sql.NullString{String: gpsLog.Sid, Valid: true},
		Lat:     gpsLog.Lat,
		Lon:     gpsLog.Lon,
		Speed:   sql.NullString{String: gpsLog.Speed, Valid: true},
		WifiLoc: gpsLog.WifiLoc,
		Battery: sql.NullInt32{Int32: int32(gpsLog.Battery), Valid: true},
	})
	if err != nil {
		logger.ErrorContext(ctx, err.Error())
		return common.NewServerError(c, common.ServerErrInternalServer)
	}
	logger.Debug("third")

	return nil
}
