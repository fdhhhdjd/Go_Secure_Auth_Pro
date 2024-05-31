package utils

import (
	"net/http"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
)

func GetXDeviceId(r *http.Request) *models.Headers {
	xDeviceId := r.Header.Get(constants.DeviceId)
	return &models.Headers{
		XDeviceId: xDeviceId,
	}
}
