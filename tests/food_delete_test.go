package tests

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoodDelete(t *testing.T) {
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

	t.Run("Delete Food Record", func(t *testing.T) {
		result := db.DB.Delete(&testFood)
		assert.NoError(t, result.Error)

		var deletedFood models.Food
		result = db.DB.First(&deletedFood, testFood.ID)
		assert.Error(t, result.Error, "Record should not be found after deletion")
	})

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testFood)
	})
}
