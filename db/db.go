package db

import (
	"OnlineFood/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:root@tcp(localhost:8889)/online_food?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	database.AutoMigrate(&models.User{},&models.Food{}, &models.Maker{}, &models.Category{})

	DB = database
}
