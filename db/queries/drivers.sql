-- name: CreateDriver :one
INSERT INTO drivers (id, name, phone, license_no, vehicle_type, vehicle_no, rating, is_available, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetDriver :one
SELECT * FROM drivers WHERE id = $1;

-- name: ListAvailableDrivers :many
SELECT * FROM drivers WHERE is_available = true;

-- name: UpdateDriverAvailability :exec
UPDATE drivers SET is_available = $2, updated_at = $3 WHERE id = $1;

-- name: UpdateDriverRating :exec
UPDATE drivers SET rating = $2, updated_at = $3 WHERE id = $1;

-- name: DeleteDriver :exec
DELETE FROM drivers WHERE id = $1;
