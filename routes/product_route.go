package routes

import (
	. "go-project/handlers"
	"go-project/middleware"
	. "go-project/repositories"
	. "go-project/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoutes(db *gorm.DB, route *gin.Engine) {

	productRepository := NewProductRepository(db)
	productService := NewProductService(productRepository)
	productHandler := NewProductHandler(productService)

	productRoute := route.Group("/api/v1/products")
	protectedRoute := productRoute.Use(middleware.JWTAuthMiddleware())
	protectedRoute.GET("/", productHandler.FindAll)
	protectedRoute.GET("/:id", productHandler.FindOne)
	protectedRoute.POST("/", productHandler.Create)
	protectedRoute.PATCH("/:id", productHandler.Update)
	protectedRoute.DELETE("/:id", productHandler.Delete)
}
