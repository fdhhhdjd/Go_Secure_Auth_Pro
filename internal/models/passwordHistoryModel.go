package models

import "database/sql"

type InsertPasswordHistoryParams struct {
	UserID       int    `json:"user_id"`
	OldPassword  string `json:"old_password"`
	ReasonStatus int    `json:"reason_status"`
}

type PasswordHistory struct {
	ID           int          `json:"id"`
	UserID       int          `json:"user_id"`
	OldPassword  string       `json:"old_password"`
	ReasonStatus int          `json:"reason_status"`
	CreatedAt    sql.NullTime `json:"created_at"`
}

// * --- Check History Password

type CheckPreviousResponse struct {
	Salt           string `json:"salt"`
	HashedPassword string `json:"hashed_password"`
}
