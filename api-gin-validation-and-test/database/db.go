package database

import (
	"api-with-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Database connection error")
	}

	DB.AutoMigrate(&models.Student{})
}
