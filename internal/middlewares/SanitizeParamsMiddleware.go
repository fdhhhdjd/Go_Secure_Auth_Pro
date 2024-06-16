package middlewares

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

// SanitizeParamsMiddleware using go-sanitize
func SanitizeParamsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := bluemonday.UGCPolicy() // User Generated Content policy
		var params = c.Request.URL.Query()

		for _, values := range params {
			for _, value := range values {
				sanitizedValue := p.Sanitize(value)
				if sanitizedValue != value {
					response.BadRequestError(c, "Potentially dangerous input detected")
					return
				}
			}
		}
		c.Next()
	}
}
