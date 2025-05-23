package tests

import (
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
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

	result := db.DB.Create(&testFood)
	assert.NoError(t, result.Error)

	updatedFood := models.Food{
		Title:      "Updated Pizza",
		Price:      19.99,
		MakerID:    2,
		CategoryID: 2,
	}

	result = db.DB.Model(&testFood).Updates(updatedFood)
	assert.NoError(t, result.Error)

	var foundFood models.Food
	result = db.DB.First(&foundFood, testFood.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, updatedFood.Title, foundFood.Title)
	assert.Equal(t, updatedFood.Price, foundFood.Price)
	assert.Equal(t, updatedFood.MakerID, foundFood.MakerID)
	assert.Equal(t, updatedFood.CategoryID, foundFood.CategoryID)

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testFood)
	})
}
