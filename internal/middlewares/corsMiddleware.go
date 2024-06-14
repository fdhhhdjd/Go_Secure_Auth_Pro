package middlewares

import (
	"net/http"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware is a middleware function that adds Cross-Origin Resource Sharing (CORS) headers to the response.
// It allows requests from different origins to access the resources of the server.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
			"https://profile-forme.com",
		}

		origin := c.Request.Header.Get("Origin")

		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == origin {
				setCORSHeaders(c.Writer, origin)

				if c.Request.Method == "OPTIONS" {
					setCORSHeaders(c.Writer, origin)
					c.Header("Access-Control-Max-Age", constants.SecondsInADay)
					c.AbortWithStatus(http.StatusOK)
					return
				}
				break
			}
		}
		c.Next()
	}
}

func setCORSHeaders(w http.ResponseWriter, origin string) {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, PUT, PATCH, DELETE, GET")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Add("Access-Control-Allow-Headers", "X-Device-Id")
	w.Header().Add("Access-Control-Allow-Headers", "X-CSRF-Token")

}
