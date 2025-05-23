package db

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_DSN")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}
	DB = database

	runMigrations()
}

func runMigrations() {
	url := os.Getenv("DB_DSN")

	m, err := migrate.New(
		"file://migrations",
		url,
	)
	if err != nil {
		log.Fatal("Ошибка создания экземпляра миграции:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Ошибка выполнения миграций:", err)
	}
}

func DbConnect() {
	dsn := "host=localhost user=r password=postgres dbname=food_db port=5432 sslmode=disable TimeZone=UTC"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}
	DB = database

}
