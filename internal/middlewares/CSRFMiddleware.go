package middlewares

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// CSRFMiddleware creates a CSRF protection middleware for Gin.
func CSRFMiddleware(secret string) gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: secret,
		ErrorFunc: func(c *gin.Context) {
			response.ForbiddenError(c)
		},
	})
}
