package database

import (
	"fmt"
	"log"

	"assignment/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		"localhost", "postgres", "mysecretpassword", "postgres", "5432",
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	database.AutoMigrate(&models.Movie{}, &models.Review{})

	DB = database
}
