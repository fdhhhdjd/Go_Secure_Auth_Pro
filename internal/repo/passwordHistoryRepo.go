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

func InsertPasswordHistory(db *sql.DB, arg models.InsertPasswordHistoryParams) error {
	_, err := db.ExecContext(context.Background(), insertPasswordHistory, arg.UserID, arg.OldPassword, arg.ReasonStatus)
	return err
}
