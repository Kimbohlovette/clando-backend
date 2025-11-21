-- name: CreateTravelHistory :one
INSERT INTO travel_histories (id, user_id, driver_id, pickup_loc, dropoff_loc, distance, fare, status, start_time, end_time, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetTravelHistory :one
SELECT * FROM travel_histories WHERE id = $1;

-- name: ListUserTravelHistory :many
SELECT * FROM travel_histories WHERE user_id = $1 ORDER BY created_at DESC;

-- name: ListDriverTravelHistory :many
SELECT * FROM travel_histories WHERE driver_id = $1 ORDER BY created_at DESC;

-- name: UpdateTravelStatus :exec
UPDATE travel_histories SET status = $2, end_time = $3 WHERE id = $1;

-- name: DeleteTravelHistory :exec
DELETE FROM travel_histories WHERE id = $1;
