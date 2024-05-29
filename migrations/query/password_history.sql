-- name: InsertPasswordHistory :exec
INSERT INTO password_history (user_id, old_password, reason_status)
VALUES ($1, $2, $3);