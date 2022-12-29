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

	r.GET("/users", controllers.FindAll)
	r.POST("users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindOne)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run(":8080")
}
