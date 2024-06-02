package routers

import (
	"os"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/utils"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/controller"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/middlewares"
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

	//* Middleware
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.SecurityHeadersMiddleware())
	r.Use(middlewares.HeadersMiddlewares())

	//* Auth
	client := r.Group("/auth")
	{
		client.GET("/veri-account", utils.AsyncHandler(controller.VerificationAccount))

		client.POST("/register", utils.AsyncHandler(controller.Register))
		client.POST("/resend-link-verification", utils.AsyncHandler(controller.ResendVerificationLink))
		client.POST("/login-identifier", utils.AsyncHandler(controller.LoginIdentifier))
		client.POST("/forget", utils.AsyncHandler(controller.ForgetPassword))

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
