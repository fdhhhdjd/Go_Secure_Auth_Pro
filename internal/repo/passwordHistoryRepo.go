package repo

import (
	"context"
	"database/sql"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
)

const insertPasswordHistory = `-- name: InsertPasswordHistory :exec
INSERT INTO password_history (user_id, old_password, reason_status)
VALUES ($1, $2, $3)
`

// InsertPasswordHistory inserts a new password history record into the database.
// It takes a database connection `db` and an `arg` parameter of type `models.InsertPasswordHistoryParams`.
// The `arg` parameter contains the necessary information for inserting the password history record.
// It returns an error if the insertion fails, otherwise it returns nil.
func InsertPasswordHistory(db *sql.DB, arg models.InsertPasswordHistoryParams) error {
	_, err := db.ExecContext(context.Background(), insertPasswordHistory, arg.UserID, arg.OldPassword, arg.ReasonStatus)
	return err
}

const checkPreviousPasswords = `-- name: CheckPreviousPasswords :one
SELECT id, user_id, old_password, reason_status, created_at
FROM password_history
WHERE user_id = $1
LIMIT $2
`

func CheckPreviousPasswords(db *sql.DB, userID int, limit int) ([]models.PasswordHistory, error) {
	rows, err := db.QueryContext(context.Background(), checkPreviousPasswords, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passwordHistories []models.PasswordHistory
	for rows.Next() {
		var i models.PasswordHistory
		err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.OldPassword,
			&i.ReasonStatus,
			&i.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		passwordHistories = append(passwordHistories, i)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return passwordHistories, nil
}
