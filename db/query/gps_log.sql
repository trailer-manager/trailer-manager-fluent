-- name: CreateGpsLog :one
INSERT INTO gps_log (
    sid,
    lat,
    lon,
    speed,
    wifi_loc,
    real_created_at
) VALUES (
     $1, $2, $3, $4, $5, now()
 ) RETURNING *;