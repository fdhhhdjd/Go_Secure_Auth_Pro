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
