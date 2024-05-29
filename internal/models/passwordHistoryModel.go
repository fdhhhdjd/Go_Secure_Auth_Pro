package models

type InsertPasswordHistoryParams struct {
	UserID       int    `json:"user_id"`
	OldPassword  string `json:"old_password"`
	ReasonStatus int    `json:"reason_status"`
}
