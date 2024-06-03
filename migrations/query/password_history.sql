-- name: InsertPasswordHistory :exec
INSERT INTO password_history (user_id, old_password, reason_status)
VALUES ($1, $2, $3);

-- name: CheckPreviousPasswords :one
SELECT *
FROM password_history
WHERE user_id = $1;