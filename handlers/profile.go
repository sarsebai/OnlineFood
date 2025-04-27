package handlers

import (
	"OnlineFood/db"
	"OnlineFood/models"

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

    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": "Неверный формат данных"})
        return
    }

    if input.Email != "" {
        user.Email = input.Email
    }

    if input.Password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
        if err != nil {
            c.JSON(500, gin.H{"error": "Ошибка при хешировании пароля"})
            return
        }
        user.Password = string(hashedPassword)
    }

    if err := db.DB.Save(&user).Error; err != nil {
        c.JSON(500, gin.H{"error": "Ошибка при обновлении профиля"})
        return
    }

    c.JSON(200, gin.H{"message": "Профиль успешно обновлен"})
} 