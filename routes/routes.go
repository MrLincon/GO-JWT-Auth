package routes

import (
	"GO-JWT-Auth/controllers"
	"GO-JWT-Auth/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to GO-JWT-Auth",
		})
	})

	authRoutes(r)
	userRoutes(r)

}

func authRoutes(r *gin.Engine) {

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/sign-up", controllers.SignUp)
		authGroup.POST("/sign-in", controllers.SignIn)
	}
}
func userRoutes(r *gin.Engine) {

	userGroup := r.Group("/user")
	{
		userGroup.GET("/fetch-user", middleware.AuthMiddleware, controllers.FetchUserDetails)

	}
}
