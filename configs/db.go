package configs

import (
	"fmt"
	. "go-project/entities"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	details := []string{"host", "port", "dbname", "user", "password"}
	var dsn string
	for _, key := range details {
		dsn += fmt.Sprintf("%v=%v ", key, os.Getenv(key))
	}
	db := connectDatabase(dsn)
	return db
}

func connectDatabase(dsn string) *gorm.DB {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&User{}, &Product{}, &Merchant{}, &Order{}, &OrderItem{})
	if err != nil {
		panic("Failed to migrate the database!")
	}

	return database
}
