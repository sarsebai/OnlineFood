package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:root@tcp(localhost:8889)/online_food?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	DB = database

	runMigrations(dsn)
}

func DbConnect() {
	dsn := "root:root@tcp(localhost:8889)/online_food?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	DB = database

}
func runMigrations(dsn string) {
	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("mysql://%s", dsn),
	)
	if err != nil {
		log.Fatal("Ошибка создания экземпляра миграции:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Ошибка выполнения миграций:", err)
	}
}
