package routers

import (
	"os"

	constants "github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/controller"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	nodeEnv := os.Getenv("ENV")

	if nodeEnv != constants.DevEnvironment {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	//* Test
	r.GET("/ping", controller.Pong)

	//* Auth
	client := r.Group("/auth")
	{
		client.POST("/register", controller.Register)
	}

	//* Not Found
	r.NoRoute(NotFound())

	//* Service Unavailable
	r.Use(ServiceUnavailable())

	return r
}

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(response.StatusNotFound, response.NotFoundError())
	}
}

func ServiceUnavailable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.AbortWithStatusJSON(response.StatusNotFound, response.ServiceUnavailable())
		}
	}

}
