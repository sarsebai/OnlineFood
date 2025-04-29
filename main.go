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
		auth.GET("/makers/:id", handlers.GetMakerByID)
		auth.POST("/makers", handlers.AddMaker)
		auth.PUT("/makers/:id", handlers.UpdateMaker)
		auth.DELETE("/makers/:id", handlers.DeleteMaker)

		auth.GET("/categories", handlers.GetCategories)
		auth.GET("/categories/:id", handlers.GetCategoryByID)
		auth.POST("/categories", handlers.AddCategory)
		auth.PUT("/categories/:id", handlers.UpdateCategory)
		auth.DELETE("/categories/:id", handlers.DeleteCategory)

		auth.GET("/profile", handlers.GetProfile)
		auth.PUT("/profile", handlers.UpdateProfile)
		auth.DELETE("/profile", handlers.DeleteProfile)
	}

	r.Run(":8088")
}
