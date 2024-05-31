package models

type Headers struct {
	XDeviceId string `json:"X-Device-Id" binding:"required"`
}
