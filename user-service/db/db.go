package db

import (
	"OnlineFood/user-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_DSN")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}
	DB = database

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Ошибка при автоматической миграции:", err)
	}
}
