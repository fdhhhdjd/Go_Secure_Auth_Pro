-- name: CreateVerification :one

INSERT INTO verification (user_id, verified_token, expires_at)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetVerification :one
SELECT * FROM verification
WHERE verified_token = $1 AND user_id = $2 LIMIT 1;

-- name: UpdateVerification :exec
UPDATE verification
SET is_verified = $1, is_active = $2
WHERE user_id = $3;

-- name: GetVerificationByUserId :one
SELECT COUNT(*) FROM verification
WHERE user_id = $1 AND is_verified = false;

-- name: UpdateVerificationBulk :exec
WITH rows_to_update AS (
    SELECT id
    FROM verification
    WHERE expires_at < NOW()
    LIMIT 50
)
UPDATE verification AS v
SET is_active = false
FROM rows_to_update AS r
WHERE v.id = r.id;
