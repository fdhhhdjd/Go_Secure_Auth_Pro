-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserId :one
SELECT * FROM users
WHERE id = $1 AND is_active = $2 LIMIT 1;


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

-- name: UpdateTwoFactorEnable :exec
UPDATE users
SET two_factor_enabled = $1
WHERE id = $2;

-- name: UpdateUser :one
UPDATE users
SET username = $1, email = $2, phone = $3, fullname = $4, hidden_email = $5, avatar = $6, gender = $7, hidden_phone_number=$8
WHERE id = $9
RETURNING id, email,username, phone,hidden_phone_number, fullname, hidden_email,avatar,gender;

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

-- name: CheckEmailExists :one
SELECT EXISTS (
    SELECT 1 
    FROM users 
    WHERE email = $1 
    AND id != $2
) AS email_exists;

-- name: UpdateEmail :exec
UPDATE users
SET email = $1
WHERE id = $2;


-- name: DestroyAccount :exec
UPDATE users
SET is_active = true
WHERE id = $1 AND is_active = false;



