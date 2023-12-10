package service

import (
	"github.com/labstack/echo/v4"
	tmError "github.com/trailer-manager/trailer-manager-common/error"
	"github.com/trailer-manager/trailer-manager-common/logger"
	db "github.com/trailer-manager/trailer-manager-fluent/db/rdb"
	model "github.com/trailer-manager/trailer-manager-fluent/model/api"

	"database/sql"
	"github.com/trailer-manager/trailer-manager-common/utility"
)

func PostGpsLog(c echo.Context) (trmErr *tmError.Error) {
	ctx := utility.GetContextFromEchoContext(c)

	gpsLog := model.GpsLogRequest{}
	if err := c.Bind(&gpsLog); err != nil {
		logger.ErrorContext(ctx, err.Error())
		return tmError.NewServerError(c, tmError.ServerErrBadRequest)
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
		return tmError.NewServerError(c, tmError.ServerErrInternalServer)
	}
	logger.Debug("third")

	return nil
}
