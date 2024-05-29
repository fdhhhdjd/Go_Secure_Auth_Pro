-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;


-- name: CreateUser :one
INSERT INTO users (
  email
) VALUES (
  $1
) RETURNING *;