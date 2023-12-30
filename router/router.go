package router

import (
	"github.com/gin-gonic/gin"
	"github.com/username/app/controllers"
	"github.com/username/app/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.POST("/login", controllers.LoginUser)
		userRoutes.PUT("/:userId", middlewares.AuthMiddleware(), controllers.UpdateUser)
		userRoutes.DELETE("/:userId", middlewares.AuthMiddleware(), controllers.DeleteUser)
	}

	photoRoutes := r.Group("/photos")
	{
		photoRoutes.POST("/", middlewares.AuthMiddleware(), controllers.CreatePhoto)
		photoRoutes.GET("/", controllers.GetPhotos)
		photoRoutes.PUT("/:photoId", middlewares.AuthMiddleware(), controllers.UpdatePhoto)
		photoRoutes.DELETE("/:photoId", middlewares.AuthMiddleware(), controllers.DeletePhoto)
	}
}
