package middlewares

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/utils"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

func HeadersMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := utils.GetXDeviceId(c.Request)

		if headers.XDeviceId == "" {
			response.BadRequestError(c)
			return
		}

		c.Set("device_id", headers.XDeviceId)
		c.Next()
	}
}
