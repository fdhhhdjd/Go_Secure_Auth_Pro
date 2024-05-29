-- name: CreateVerification :one

INSERT INTO verification (user_id, verified_token, expires_at)
VALUES ($1, $2, $3) RETURNING *;