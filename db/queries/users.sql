-- name: CreateUser :one
INSERT INTO users (id, username, phone)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByPhone :one
SELECT * FROM users WHERE phone = $1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: UpdateUser :one
UPDATE users SET username = $2, phone = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
