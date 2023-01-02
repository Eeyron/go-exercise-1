package main

import (
	"go-project/controllers"
	"go-project/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.Init()

	r.GET("/users", controllers.FindAllUsers)
	r.POST("users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindOneUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.GET("/merchants", controllers.FindAllMerchants)
	r.POST("merchants", controllers.CreateMerchant)
	r.GET("/merchants/:id", controllers.FindOneMerchant)
	r.PATCH("/merchants/:id", controllers.UpdateMerchant)
	r.DELETE("/merchants/:id", controllers.DeleteMerchant)

	r.Run(":8080")
}
