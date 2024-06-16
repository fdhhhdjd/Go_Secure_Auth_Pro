package middlewares

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// IPBlackList is a middleware function that checks if the client's IP address is blacklisted.
// If the IP address is found in the blacklist, it returns a forbidden error response.
// Otherwise, it allows the request to proceed to the next middleware or handler.
func IPBlackList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		// Check if the IP is in the blacklist
		isBlacklisted, err := global.Cache.SIsMember(c, constants.BlackListIP, ip).Result()
		if err != nil {
			response.InternalServerError(c)
			return
		}

		if isBlacklisted {
			response.ForbiddenError(c)
			return
		}

		c.Next()
	}
}
