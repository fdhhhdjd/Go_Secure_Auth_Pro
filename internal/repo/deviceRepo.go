package repo

import (
	"context"
	"database/sql"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
)

// upsetDevice represents the SQL query for upserting a device record in the devices table.
// It inserts a new record if the device_id does not exist, otherwise it updates the existing record.
// The query returns the id, user_id, device_id, device_type, logged_in_at, logged_out_at, ip, public_key,
// is_active, created_at, and updated_at fields of the inserted/updated record.
const upsetDevice = `-- name: UpsertDevice :one
INSERT INTO devices (
  user_id, 
  device_id, 
  device_type, 
  logged_in_at, 
  logged_out_at,
  ip, 
  public_key, 
  is_active,
  created_at, 
  updated_at
) VALUES (
  $1, $2, $3, CURRENT_TIMESTAMP, NULL, $4, $5, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
)
ON CONFLICT (device_id) DO UPDATE SET
  user_id = excluded.user_id,
  device_type = excluded.device_type,
  logged_in_at = excluded.logged_in_at,
  logged_out_at = excluded.logged_out_at,
  ip = excluded.ip,
  public_key = excluded.public_key,
  is_active = excluded.is_active,
  updated_at = excluded.updated_at
RETURNING id, user_id, device_id, device_type, logged_in_at, logged_out_at, ip, public_key, is_active, created_at, updated_at
`

func UpsetDevice(db *sql.DB, arg models.UpsetDeviceParams) (models.Device, error) {
	row := db.QueryRowContext(context.Background(), upsetDevice,
		arg.UserID,
		arg.DeviceID,
		arg.DeviceType,
		arg.Ip,
		arg.PublicKey,
	)
	var i models.Device
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceID,
		&i.DeviceType,
		&i.LoggedInAt,
		&i.LoggedOutAt,
		&i.Ip,
		&i.PublicKey,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDeviceId = `-- name: GetDeviceId :one
SELECT id, user_id, device_id, device_type, logged_in_at, logged_out_at, ip, public_key, is_active, created_at, updated_at FROM devices
WHERE device_id = $1 AND is_active = $2 LIMIT 1
`

func GetDeviceId(db *sql.DB, arg models.GetDeviceIdParams) (models.Device, error) {
	row := db.QueryRowContext(context.Background(), getDeviceId, arg.DeviceId, arg.IsActive)
	var i models.Device
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceID,
		&i.DeviceType,
		&i.LoggedInAt,
		&i.LoggedOutAt,
		&i.Ip,
		&i.PublicKey,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
