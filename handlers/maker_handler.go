package handlers

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMakers(c *gin.Context) {
	var makers []models.Maker
	db.DB.Find(&makers)
	c.JSON(http.StatusOK, makers)
}

func GetMakerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maker models.Maker
	result := db.DB.First(&maker, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Производитель не найден"})
		return
	}
	c.JSON(http.StatusOK, maker)
}

func AddMaker(c *gin.Context) {
	var newMaker models.Maker
	if err := c.ShouldBindJSON(&newMaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&newMaker)
	c.JSON(http.StatusCreated, newMaker)
}

func UpdateMaker(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maker models.Maker
	if err := db.DB.First(&maker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Производитель не найден"})
		return
	}
	if err := c.ShouldBindJSON(&maker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	db.DB.Save(&maker)
	c.JSON(http.StatusOK, maker)
}

func DeleteMaker(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maker models.Maker
	if err := db.DB.First(&maker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Производитель не найден"})
		return
	}
	db.DB.Delete(&maker)
	c.JSON(http.StatusOK, gin.H{"message": "Производитель успешно удален"})
}
