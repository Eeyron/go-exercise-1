package main

import (
	"go-project/controllers"
	"go-project/models"
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

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)

	r.Run()
}
