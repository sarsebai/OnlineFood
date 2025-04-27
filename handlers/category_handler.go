package handlers

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	db.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func AddCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&newCategory)
	c.JSON(http.StatusCreated, newCategory)
}
