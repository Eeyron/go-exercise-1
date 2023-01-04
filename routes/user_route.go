package routes

import (
	. "go-project/handlers"
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
	userRoute.GET("/", userHandler.FindAll)
	userRoute.GET("/:id", userHandler.FindOne)
	userRoute.POST("/", userHandler.Create)
	userRoute.PATCH("/:id", userHandler.Update)
	userRoute.DELETE("/:id", userHandler.Delete)
}
