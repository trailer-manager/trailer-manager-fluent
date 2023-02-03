// Code generated by sqlc. DO NOT EDIT.
// source: sensor.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createSensor = `-- name: CreateSensor :one
INSERT INTO sensor (
    sid
) VALUES (
	 $1
 ) RETURNING sid, uid, lat, lon, wifi_loc, created_at, updated_at
`

func (q *Queries) CreateSensor(ctx context.Context, sid string) (Sensor, error) {
	row := q.db.QueryRowContext(ctx, createSensor, sid)
	var i Sensor
	err := row.Scan(
		&i.Sid,
		&i.Uid,
		&i.Lat,
		&i.Lon,
		pq.Array(&i.WifiLoc),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSensor = `-- name: GetSensor :one
SELECT sid, uid, lat, lon, wifi_loc, created_at, updated_at FROM sensor
WHERE sid = $1 LIMIT 1
`

func (q *Queries) GetSensor(ctx context.Context, sid string) (Sensor, error) {
	row := q.db.QueryRowContext(ctx, getSensor, sid)
	var i Sensor
	err := row.Scan(
		&i.Sid,
		&i.Uid,
		&i.Lat,
		&i.Lon,
		pq.Array(&i.WifiLoc),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSensorLocation = `-- name: UpdateSensorLocation :one
UPDATE sensor
SET lat = $2,
    lon = $3,
    wifi_loc = $4,
    updated_at = now()
WHERE sid = $1
    RETURNING sid, uid, lat, lon, wifi_loc, created_at, updated_at
`

type UpdateSensorLocationParams struct {
	Sid     string         `json:"sid"`
	Lat     sql.NullString `json:"lat"`
	Lon     sql.NullString `json:"lon"`
	WifiLoc []string       `json:"wifi_loc"`
}

func (q *Queries) UpdateSensorLocation(ctx context.Context, arg UpdateSensorLocationParams) (Sensor, error) {
	row := q.db.QueryRowContext(ctx, updateSensorLocation,
		arg.Sid,
		arg.Lat,
		arg.Lon,
		pq.Array(arg.WifiLoc),
	)
	var i Sensor
	err := row.Scan(
		&i.Sid,
		&i.Uid,
		&i.Lat,
		&i.Lon,
		pq.Array(&i.WifiLoc),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
