package routes

import (
	. "go-project/handlers"
	"go-project/middleware"
	. "go-project/repositories"
	. "go-project/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, route *gin.Engine) {

	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	userRoute := route.Group("/api/v1/users")
	userRoute.POST("/", userHandler.Create)

	protectedRoute := userRoute.Use(middleware.JWTAuthMiddleware())
	protectedRoute.GET("/", userHandler.FindAll)
	protectedRoute.GET("/:id", userHandler.FindOne)
	protectedRoute.PATCH("/:id", userHandler.Update)
	protectedRoute.DELETE("/:id", userHandler.Delete)
}
