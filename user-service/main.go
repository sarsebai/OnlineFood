package main

import (
	"OnlineFood/auth"
	"OnlineFood/user-service/db"
	"OnlineFood/user-service/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.New()
	r.Use(auth.LoggerMiddleware())
	r.Use(gin.Recovery())

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	profile := r.Group("/profile")
	profile.Use(auth.AuthMiddleware())
	{
		profile.GET("", handlers.GetProfile)
		profile.PUT("", handlers.UpdateProfile)
		profile.DELETE("", handlers.DeleteProfile)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
