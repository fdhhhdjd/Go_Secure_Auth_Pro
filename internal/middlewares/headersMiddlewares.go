package middlewares

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/utils"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// HeadersMiddlewares is a middleware function that validates the X-Device-Id header in the request.
// If the X-Device-Id header is missing or empty, it responds with a bad request error.
// Otherwise, it sets the "device_id" value in the context and proceeds to the next middleware or handler.
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
