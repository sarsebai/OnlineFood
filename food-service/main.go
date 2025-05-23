package main

import (
	"OnlineFood/auth"
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.New()
	r.Use(auth.LoggerMiddleware())
	r.Use(gin.Recovery())

	userHandler := handlers.NewUserHandler()
	r.GET("/profile", auth.AuthMiddleware(), userHandler.GetProfile)

	foodGroup := r.Group("/foods")
	foodGroup.Use(auth.AuthMiddleware())
	{
		foodGroup.POST("/", handlers.AddFood)
		foodGroup.GET("/", handlers.GetFoods)
		foodGroup.GET("/:id", handlers.GetFoodByID)
		foodGroup.PUT("/:id", handlers.UpdateFood)
		foodGroup.DELETE("/:id", handlers.DeleteFood)
	}

	makerGroup := r.Group("/makers")
	makerGroup.Use(auth.AuthMiddleware())
	{
		makerGroup.POST("/", handlers.AddMaker)
		makerGroup.GET("/", handlers.GetMakers)
		makerGroup.GET("/:id", handlers.GetMakerByID)
		makerGroup.PUT("/:id", handlers.UpdateMaker)
		makerGroup.DELETE("/:id", handlers.DeleteMaker)
	}

	categoryGroup := r.Group("/categories")
	categoryGroup.Use(auth.AuthMiddleware())
	{
		categoryGroup.POST("/", handlers.AddCategory)
		categoryGroup.GET("/", handlers.GetCategories)
		categoryGroup.GET("/:id", handlers.GetCategoryByID)
		categoryGroup.PUT("/:id", handlers.UpdateCategory)
		categoryGroup.DELETE("/:id", handlers.DeleteCategory)
	}

	println("FOOD SERVICE ON PORT 8090")

	r.Run(":8090")
}
