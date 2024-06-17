package middlewares

import (
	"path/filepath"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// PathTraversalMiddleware is a middleware function that checks for path traversal attacks.
// It ensures that the requested path is valid and does not contain any path traversal characters.
// If a path traversal attack is detected, it returns a BadRequestError response.
func PathTraversalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		cleanPath := filepath.Clean(path)

		if cleanPath != path {
			response.BadRequestError(c, response.ErrPathTraversal)
			return
		}

		c.Next()
	}
}
