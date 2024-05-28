package utils

import "github.com/gin-gonic/gin"

func AsyncHandler(fn func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(c); err != nil {
			c.Error(err)
			c.Next()
		}
	}
}
