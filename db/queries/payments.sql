-- name: CreatePayment :one
INSERT INTO payments (id, user_id, travel_id, amount, status, payment_method, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments WHERE id = $1;

-- name: GetPaymentByTravelID :one
SELECT * FROM payments WHERE travel_id = $1;

-- name: ListUserPayments :many
SELECT * FROM payments WHERE user_id = $1 ORDER BY created_at DESC;

-- name: UpdatePaymentStatus :exec
UPDATE payments SET status = $2 WHERE id = $1;

-- name: DeletePayment :exec
DELETE FROM payments WHERE id = $1;
