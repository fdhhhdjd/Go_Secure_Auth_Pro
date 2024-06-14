package routers

import (
	"os"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/utils"
	_ "github.com/fdhhhdjd/Go_Secure_Auth_Pro/docs/swagger"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/controller"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/middlewares"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {

	nodeEnv := os.Getenv("ENV")

	if nodeEnv != constants.DevEnvironment {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	//* Swaggers
	r.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//* Test
	r.GET("/ping", controller.Pong)

	//* Middleware
	// CSRF middleware
	secret := os.Getenv("CSRF_TOKEN") // Replace with your actual secret
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions(constants.CSRFToken, store))
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.SecurityHeadersMiddleware())
	r.Use(middlewares.HeadersMiddlewares())
	r.Use(middlewares.CSRFMiddleware(secret))
	r.Use(middlewares.RateLimiter(5, 10)) // 5 requests per second, with a burst of 10

	//* Group v1 routes
	v1 := r.Group("/v1")
	{
		//* Group v1/key routes
		key := v1.Group("/key")
		{
			key.GET("/csrf-token", utils.AsyncHandler(controller.GetCsRfToken))
		}

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
			auth.POST("/verify-otp", utils.AsyncHandler(controller.VerificationOtp))

			auth.Use(middlewares.RefetchTokenMiddleware())
			auth.GET("/renew-token", utils.AsyncHandler(controller.RenewToken))

		}

		//* Group v1/user routes (example, you can add more routes here)
		user := v1.Group("/user")
		{
			user.Use(middlewares.AuthorizationMiddleware())

			user.GET("/logout", utils.AsyncHandler(controller.LogoutUser))
			user.GET("/profile/:id", utils.AsyncHandler(controller.GetProfileUser))
			user.GET("/destroy-account", utils.AsyncHandler(controller.DestroyAccount))

			user.POST("/update-profile", utils.AsyncHandler(controller.UpdateProfile))
			user.POST("/change-pass", utils.AsyncHandler(controller.ChangePassword))
			user.POST("/enable-tow-factor", utils.AsyncHandler(controller.EnableTowFactor))
			user.POST("/send-otp-update-email", utils.AsyncHandler(controller.SendOtpUpdateEmail))
			user.POST("/update-email", utils.AsyncHandler(controller.UpdateEmailUser))

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
