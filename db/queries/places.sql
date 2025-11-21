-- name: CreatePlace :one
INSERT INTO places (id, name, address, latitude, longitude, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPlace :one
SELECT * FROM places WHERE id = $1;

-- name: ListPlaces :many
SELECT * FROM places ORDER BY name;

-- name: DeletePlace :exec
DELETE FROM places WHERE id = $1;
