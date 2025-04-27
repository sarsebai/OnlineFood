package main

import (
	"OnlineFood/db"
	"OnlineFood/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(handlers.AuthMiddleware())
	{
		auth.GET("/foods", handlers.GetFoods)
		auth.GET("/foods/:id", handlers.GetFoodByID)
		auth.POST("/foods", handlers.AddFood)
		auth.PUT("/foods/:id", handlers.UpdateFood)
		auth.DELETE("/foods/:id", handlers.DeleteFood)

		auth.GET("/makers", handlers.GetMakers)
		auth.POST("/makers", handlers.AddMaker)

		auth.GET("/categories", handlers.GetCategories)
		auth.POST("/categories", handlers.AddCategory)

		auth.GET("/profile", handlers.GetProfile)
		auth.PUT("/profile", handlers.UpdateProfile)
	}

	r.Run(":8088")
}
