package handlers

import (
	"OnlineFood/user-service/db"
	"OnlineFood/user-service/models"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetProfile(c *gin.Context) {
	username := c.GetString("username")
	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(200, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
	})
}

func UpdateProfile(c *gin.Context) {
	username := c.GetString("username")
	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
		return
	}
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Ошибка хеширования пароля при обновлении: %v", err) // Логируем ошибку хеширования
			c.JSON(500, gin.H{"error": "Ошибка при хешировании пароля"})
			return
		}
		user.Password = string(hashedPassword)
	}
	if err := db.DB.Save(&user).Error; err != nil {
		log.Printf("Ошибка базы данных при обновлении профиля: %v", err) // Логируем ошибку базы данных
		c.JSON(500, gin.H{"error": "Ошибка при обновлении профиля"})
		return
	}

	c.JSON(200, gin.H{"message": "Профиль успешно обновлен"})
}

func DeleteProfile(c *gin.Context) {
	username := c.GetString("username")
	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
		return
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Ошибка при удалении профиля"})
		return
	}

	c.JSON(200, gin.H{"message": "Профиль успешно удален"})
}
