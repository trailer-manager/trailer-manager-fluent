// Code generated by sqlc. DO NOT EDIT.
// source: gps_log.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createGpsLog = `-- name: CreateGpsLog :one
INSERT INTO gps_log (
    sid,
    lat,
    lon,
    speed,
    wifi_loc,
    battery,
    real_created_at
) VALUES (
     $1, $2, $3, $4, $5, $6, now()
 ) RETURNING seq, sid, lat, lon, speed, wifi_loc, battery, real_created_at
`

type CreateGpsLogParams struct {
	Sid     sql.NullString `json:"sid"`
	Lat     string         `json:"lat"`
	Lon     string         `json:"lon"`
	Speed   sql.NullString `json:"speed"`
	WifiLoc []string       `json:"wifi_loc"`
	Battery sql.NullInt32  `json:"battery"`
}

func (q *Queries) CreateGpsLog(ctx context.Context, arg CreateGpsLogParams) (GpsLog, error) {
	row := q.db.QueryRowContext(ctx, createGpsLog,
		arg.Sid,
		arg.Lat,
		arg.Lon,
		arg.Speed,
		pq.Array(arg.WifiLoc),
		arg.Battery,
	)
	var i GpsLog
	err := row.Scan(
		&i.Seq,
		&i.Sid,
		&i.Lat,
		&i.Lon,
		&i.Speed,
		pq.Array(&i.WifiLoc),
		&i.Battery,
		&i.RealCreatedAt,
	)
	return i, err
}
