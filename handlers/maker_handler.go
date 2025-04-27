package handlers

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMakers(c *gin.Context) {
	var makers []models.Maker
	db.DB.Find(&makers)
	c.JSON(http.StatusOK, makers)
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
