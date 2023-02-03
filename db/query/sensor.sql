-- name: CreateSensor :one
INSERT INTO sensor (
    sid
) VALUES (
     $1
 ) RETURNING *;

-- name: GetSensor :one
SELECT * FROM sensor
WHERE sid = $1 LIMIT 1;


-- name: UpdateSensorLocation :one
UPDATE sensor
SET lat = $2,
    lon = $3,
    wifi_loc = $4,
    updated_at = now()
WHERE sid = $1
    RETURNING *;