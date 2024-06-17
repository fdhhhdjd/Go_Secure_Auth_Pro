package middlewares

import (
	"log"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// ContentTypeValidationMiddleware is a middleware function that validates the content type of incoming requests.
// It checks if the content type is "application/json" and returns an unsupported media type error if not.
// It then calls the next handler in the chain.
func ContentTypeValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			contentType := c.Request.Header.Get("Content-Type")
			log.Println(contentType)
			if contentType != "application/json" {
				response.UnSupportMediaTypeError(c, response.ErrCodeContentType)
				return
			}
		}
		c.Next()
	}
}
