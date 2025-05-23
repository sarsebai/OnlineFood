package handlers

import (
	"OnlineFood/food-service/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	baseURL    string
	httpClient *http.Client
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		baseURL: "http://user-service:8080",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не предоставлен"})
		return
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/profile", h.baseURL), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка создания запроса: %v", err)})
		return
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.httpClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при запросе к user-service: %v", err)})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Неверный статус ответа: %d", resp.StatusCode)})
		return
	}

	var profile models.UserProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при обработке ответа: %v", err)})
		return
	}

	c.JSON(http.StatusOK, profile)
}
