package routes

import (
	. "go-project/handlers"
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
	merchantRoute.GET("/", merchantHandler.FindAll)
	merchantRoute.GET("/:id", merchantHandler.FindOne)
	merchantRoute.POST("/", merchantHandler.Create)
	merchantRoute.PATCH("/:id", merchantHandler.Update)
	merchantRoute.DELETE("/:id", merchantHandler.Delete)
}
