package handlers

import (
	"OnlineFood/auth"
	"OnlineFood/user-service/db"
	"OnlineFood/user-service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Ошибка привязки JSON: %v", err)
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		c.JSON(400, gin.H{"error": "Все поля обязательны для заполнения"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Ошибка хеширования пароля: %v", err)
		c.JSON(500, gin.H{"error": "Ошибка при хешировании пароля"})
		return
	}
	user.Password = string(hashedPassword)

	if user.Role == "" {
		user.Role = "user"
	}

	if err := db.DB.Create(&user).Error; err != nil {
		log.Printf("Ошибка создания пользователя: %v", err)
		if err.Error() == "duplicate key value violates unique constraint" {
			c.JSON(400, gin.H{"error": "Пользователь с таким именем или email уже существует"})
		} else {
			c.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при создании пользователя: %v", err)})
		}
		return
	}

	c.JSON(200, gin.H{
		"message": "Пользователь успешно создан",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	var user models.User
	if err := db.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Неверное имя пользователя или пароль"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Неверное имя пользователя или пароль"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
	})
	tokenString, err := token.SignedString(auth.JWTKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка при создании токена"})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
		"user": gin.H{
			"username": user.Username,
			"role":     user.Role,
		},
	})
}
