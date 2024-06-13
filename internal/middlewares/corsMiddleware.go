package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware is a middleware function that adds Cross-Origin Resource Sharing (CORS) headers to the response.
// It allows requests from different origins to access the resources of the server.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := []string{
			"http://localhost:5173",
		}

		origin := c.Request.Header.Get("Origin")

		var isAllowedOrigin bool
		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == origin {
				isAllowedOrigin = true
				break
			}
		}

		if !isAllowedOrigin {
			c.Next()
			return
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, PUT, PATCH, DELETE, GET")

		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, PUT, PATCH, DELETE, GET")
			c.Header("Access-Control-Max-Age", "86400")
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
