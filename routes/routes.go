package routes

import (
	"GO-JWT-Auth/controllers"
	"GO-JWT-Auth/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SetupRoutes initializes the routes for the application.
func SetupRoutes(r *gin.Engine, logger *zap.SugaredLogger) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to GO-JWT-Auth",
		})
	})

	authRoutes(r, logger)
	userRoutes(r, logger)
}

// authRoutes sets up the authentication-related routes.
func authRoutes(r *gin.Engine, logger *zap.SugaredLogger) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/sign-up", func(c *gin.Context) {
			controllers.SignUp(c, logger)
		})
		authGroup.POST("/sign-in", func(c *gin.Context) {
			controllers.SignIn(c, logger)
		})
		authGroup.POST("/send-otp", func(c *gin.Context) {
			controllers.SendOtp(c, logger)
		})
		authGroup.POST("/reset-password", func(c *gin.Context) {
			controllers.ResetPassword(c, logger)
		})
	}
}

// userRoutes sets up the user-related routes.
func userRoutes(r *gin.Engine, logger *zap.SugaredLogger) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/fetch-user", middleware.AuthMiddleware, func(c *gin.Context) {
			controllers.FetchUserDetails(c, logger)
		})
	}
}
