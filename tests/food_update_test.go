package tests

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoodUpdate(t *testing.T) {
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

	t.Run("Update Food Record", func(t *testing.T) {
		newPrice := 17.99
		result := db.DB.Model(&testFood).Update("price", newPrice)
		assert.NoError(t, result.Error)

		var updatedFood models.Food
		db.DB.First(&updatedFood, testFood.ID)
		assert.Equal(t, newPrice, updatedFood.Price)
	})

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testFood)
	})
}
