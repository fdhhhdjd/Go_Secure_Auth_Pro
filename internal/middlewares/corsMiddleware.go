package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware is a middleware function that adds Cross-Origin Resource Sharing (CORS) headers to the response.
// It allows requests from different origins to access the resources of the server.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, PUT, PATCH, DELETE, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}

		if c.Request.Method == "PUT" {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}

		c.Next()
	}
}
