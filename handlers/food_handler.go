package handlers

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFoods(c *gin.Context) {
	var foods []models.Food
	db.DB.Find(&foods)
	c.JSON(http.StatusOK, foods)
}

func GetFoodByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var food models.Food
	result := db.DB.First(&food, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Блюдо не найдено"})
		return
	}
	c.JSON(http.StatusOK, food)
}

func AddFood(c *gin.Context) {
	var newFood models.Food
	if err := c.ShouldBindJSON(&newFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	db.DB.Create(&newFood)
	c.JSON(http.StatusCreated, newFood)
}

func UpdateFood(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var food models.Food
	if err := db.DB.First(&food, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Блюдо не найдено"})
		return
	}
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	db.DB.Save(&food)
	c.JSON(http.StatusOK, food)
}

func DeleteFood(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var food models.Food
	if err := db.DB.First(&food, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Блюдо не найдено"})
		return
	}
	db.DB.Delete(&food)
	c.JSON(http.StatusOK, gin.H{"message": "Блюдо успешно удалено"})
}
