package models

import (
	"database/sql"
	"time"
)

type Device struct {
	ID          int            `json:"id"`
	UserID      int            `json:"user_id"`
	DeviceID    string         `json:"device_id"`
	DeviceType  string         `json:"device_type"`
	LoggedInAt  time.Time      `json:"logged_in_at"`
	LoggedOutAt sql.NullTime   `json:"logged_out_at"`
	Ip          sql.NullString `json:"ip"`
	PublicKey   sql.NullString `json:"public_key"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type Headers struct {
	XDeviceId string `json:"X-Device-Id" binding:"required"`
}

type UpsetDeviceParams struct {
	UserID     int            `json:"user_id"`
	DeviceID   string         `json:"device_id"`
	Ip         sql.NullString `json:"ip"`
	DeviceType string         `json:"device_type"`
	PublicKey  string         `json:"public_key"`
}

// * -- Get Device
type GetDeviceIdParams struct {
	DeviceId string `json:"device_id"`
	IsActive bool   `json:"is_active"`
}

// * --- Update Time Logout
type UpdateTimeLogoutParams struct {
	LoggedOutAt sql.NullTime `json:"logged_out_at"`
	DeviceId    string       `json:"device_id"`
}
