package middlewares

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter returns a gin.HandlerFunc that limits the rate of requests
func RateLimiter(rps float64, burst int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(rps), burst)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			response.TooManyRequestsError(c, response.ErrCodeTooManyRequests)
			return
		}
		c.Next()
	}
}
