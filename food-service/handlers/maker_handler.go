package handlers

import (
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMakers(c *gin.Context) {
	var makers []models.Maker
	if err := db.DB.Find(&makers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка производителей"})
		return
	}
	c.JSON(http.StatusOK, makers)
}

func GetMakerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var maker models.Maker
	if err := db.DB.First(&maker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Производитель не найден"})
		return
	}
	c.JSON(http.StatusOK, maker)
}

func AddMaker(c *gin.Context) {
	var newMaker models.Maker
	if err := c.ShouldBindJSON(&newMaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := db.DB.Create(&newMaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании производителя"})
		return
	}

	c.JSON(http.StatusCreated, newMaker)
}

func UpdateMaker(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var maker models.Maker
	if err := db.DB.First(&maker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Производитель не найден"})
		return
	}

	if err := c.ShouldBindJSON(&maker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := db.DB.Save(&maker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении производителя"})
		return
	}

	c.JSON(http.StatusOK, maker)
}

func DeleteMaker(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var maker models.Maker
	if err := db.DB.First(&maker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Производитель не найден"})
		return
	}

	if err := db.DB.Delete(&maker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении производителя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Производитель успешно удален"})
}
