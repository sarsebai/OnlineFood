package handlers

import (
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFoods(c *gin.Context) {
	var foods []models.Food
	if err := db.DB.Find(&foods).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка блюд"})
		return
	}
	c.JSON(http.StatusOK, foods)
}

func GetFoodByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var food models.Food
	if err := db.DB.First(&food, id).Error; err != nil {
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

	if err := db.DB.Create(&newFood).Error; err != nil {
		log.Printf("Ошибка базы данных при создании блюда: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании блюда"})
		return
	}

	c.JSON(http.StatusCreated, newFood)
}

func UpdateFood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var food models.Food
	if err := db.DB.First(&food, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Блюдо не найдено"})
		return
	}

	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := db.DB.Save(&food).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении блюда"})
		return
	}

	c.JSON(http.StatusOK, food)
}

func DeleteFood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var food models.Food
	if err := db.DB.First(&food, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Блюдо не найдено"})
		return
	}

	if err := db.DB.Delete(&food).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении блюда"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Блюдо успешно удалено"})
}
