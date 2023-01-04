package routes

import (
	. "go-project/handlers"
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
	productRoute.GET("/", productHandler.FindAll)
	productRoute.GET("/:id", productHandler.FindOne)
	productRoute.POST("/", productHandler.Create)
	productRoute.PATCH("/:id", productHandler.Update)
	productRoute.DELETE("/:id", productHandler.Delete)
}
