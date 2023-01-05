package routes

import (
	. "go-project/handlers"
	"go-project/middleware"
	. "go-project/repositories"
	. "go-project/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MerchantRoutes(db *gorm.DB, route *gin.Engine) {

	merchantRepository := NewMerchantRepository(db)
	merchantService := NewMerchantService(merchantRepository)
	merchantHandler := NewMerchantHandler(merchantService)

	merchantRoute := route.Group("/api/v1/merchants")
	protectedRoute := merchantRoute.Use(middleware.JWTAuthMiddleware())
	protectedRoute.GET("/", merchantHandler.FindAll)
	protectedRoute.GET("/:id", merchantHandler.FindOne)
	protectedRoute.POST("/", merchantHandler.Create)
	protectedRoute.PATCH("/:id", merchantHandler.Update)
	protectedRoute.DELETE("/:id", merchantHandler.Delete)
}
