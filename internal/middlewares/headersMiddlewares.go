package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HeadersMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("OK")
		c.Next()
	}
}
