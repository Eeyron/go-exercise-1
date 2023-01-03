package db

import (
	"fmt"
	. "go-project/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	details := []string{"host", "port", "dbname", "user", "password"}
	var dsn string
	for _, key := range details {
		dsn += fmt.Sprintf("%v=%v ", key, os.Getenv(key))
	}
	connectDatabase(dsn)
}

func connectDatabase(dsn string) {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&User{}, &Product{}, &Merchant{}, &Order{}, &OrderItem{})
	if err != nil {
		panic("Failed to migrate the database!")
	}

	DB = database
}
