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

-- name: JoinUsersWithVerificationByEmail :many
SELECT users.*
FROM users
JOIN verification ON users.id = verification.user_id
WHERE verification.is_verified = true
AND users.email = $1;

-- name: JoinUsersWithVerificationByPhone :many
SELECT users.*
FROM users
JOIN verification ON users.id = verification.user_id
WHERE verification.is_verified = true
AND users.phone = $1;

-- name: JoinUsersWithVerificationByUsername :many
SELECT users.*
FROM users
JOIN verification ON users.id = verification.user_id
WHERE verification.is_verified = true
AND users.username = $1;