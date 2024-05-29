-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;


-- name: CreateUser :one
INSERT INTO users (
  email
) VALUES (
  $1
) RETURNING *;

-- name: UpdatePassword :exec
UPDATE users
SET password_hash = $1
WHERE id = $2;