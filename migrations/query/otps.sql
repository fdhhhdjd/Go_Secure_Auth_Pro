-- name: CreateOtp :one
INSERT INTO otps (
    user_id,
    otp_code,
    expires_at
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetNewOtps :many
SELECT otps.*, users.email
FROM otps
JOIN users ON otps.user_id = users.id
WHERE otps.expires_at > NOW()
    AND otps.otp_code = $1
    AND otps.is_active = TRUE;

-- name: UpdateOtpIsActive :exec
UPDATE otps
SET is_active = $1
WHERE otp_code = $2;