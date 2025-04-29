package tests

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoodPersistence(t *testing.T) {
	db.DbConnect()

	testFood := models.Food{
		Title:      "Test Pizza",
		Price:      15.99,
		MakerID:    1,
		CategoryID: 1,
	}

	t.Run("Create Food Record", func(t *testing.T) {
		result := db.DB.Create(&testFood)
		assert.NoError(t, result.Error)
		assert.NotZero(t, testFood.ID, "ID should not be zero after creation")
	})

	t.Run("Read Food Record", func(t *testing.T) {
		var foundFood models.Food
		result := db.DB.First(&foundFood, testFood.ID)
		assert.NoError(t, result.Error)
		assert.Equal(t, testFood.Title, foundFood.Title)
		assert.Equal(t, testFood.Price, foundFood.Price)
		assert.Equal(t, testFood.MakerID, foundFood.MakerID)
		assert.Equal(t, testFood.CategoryID, foundFood.CategoryID)
	})

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testFood)
	})
}
