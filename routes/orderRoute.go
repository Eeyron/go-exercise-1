package routes

import (
	. "go-project/handlers"
	"go-project/middleware"
	. "go-project/repositories"
	. "go-project/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRoutes(db *gorm.DB, route *gin.Engine) {

	orderRepository := NewOrderRepository(db)
	orderService := NewOrderService(orderRepository)
	orderHandler := NewOrderHandler(orderService)

	orderRoute := route.Group("/api/v1/orders")
	protectedRoute := orderRoute.Use(middleware.JWTAuthMiddleware())
	protectedRoute.GET("/", orderHandler.FindAll)
	protectedRoute.GET("/:id", orderHandler.FindOne)
	protectedRoute.POST("/", orderHandler.Create)
}
