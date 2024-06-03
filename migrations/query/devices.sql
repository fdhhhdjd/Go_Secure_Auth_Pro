-- name: UpsertDevice :one
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
RETURNING *;


-- name: GetDeviceId :one
SELECT * FROM devices
WHERE device_id = $1 AND is_active = $2 LIMIT 1;

-- name: UpdateTimeLogout :exec
UPDATE devices
SET logged_out_at = $1
WHERE id = $2;