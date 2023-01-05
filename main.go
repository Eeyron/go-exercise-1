package main

import (
	config "go-project/configs"
	route "go-project/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := config.SetupDatabase()

	router := SetupRouter(db)

	router.Run(":" + os.Getenv("SERVER"))
}

func SetupRouter(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	route.UserRoutes(db, router)
	route.MerchantRoutes(db, router)
	route.ProductRoutes(db, router)
	route.OrderRoutes(db, router)

	return router
}
