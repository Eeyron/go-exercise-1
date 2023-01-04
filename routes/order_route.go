package routes

import (
	. "go-project/handlers"
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
	orderRoute.GET("/", orderHandler.FindAll)
	orderRoute.GET("/:id", orderHandler.FindOne)
	orderRoute.POST("/", orderHandler.Create)
}
