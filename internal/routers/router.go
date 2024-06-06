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

	//* Group v1 routes
	v1 := r.Group("/v1")
	{
		//* Group v1/auth routes
		auth := v1.Group("/auth")
		{
			auth.GET("/veri-account", utils.AsyncHandler(controller.VerificationAccount))

			auth.POST("/register", utils.AsyncHandler(controller.Register))
			auth.POST("/resend-link-verification", utils.AsyncHandler(controller.ResendVerificationLink))
			auth.POST("/login-identifier", utils.AsyncHandler(controller.LoginIdentifier))
			auth.POST("/login-social", utils.AsyncHandler(controller.LoginSocial))
			auth.POST("/forget", utils.AsyncHandler(controller.ForgetPassword))
			auth.POST("/reset-password", utils.AsyncHandler(controller.ResetPassword))

			auth.Use(middlewares.RefetchTokenMiddleware())
			auth.GET("/renew-token", utils.AsyncHandler(controller.RenewToken))

		}

		//* Group v1/user routes (example, you can add more routes here)
		user := v1.Group("/user")
		{
			user.Use(middlewares.AuthorizationMiddleware())

			user.GET("/logout", utils.AsyncHandler(controller.LogoutUser))
			user.GET("/profile/:id", utils.AsyncHandler(controller.GetProfileUser))

			user.POST("/update-profile", utils.AsyncHandler(controller.UpdateProfile))
			user.POST("/change-pass", utils.AsyncHandler(controller.ChangePassword))
			user.POST("/enable-tow-factor", utils.AsyncHandler(controller.EnableTowFactor))
		}
	}

	//* Not Found
	r.NoRoute(NotFound())

	//* Service Unavailable
	r.Use(ServiceUnavailable())

	return r
}

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.NotFoundError(c)
	}
}

func ServiceUnavailable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			response.ServiceUnavailable(c)
		}
	}

}
