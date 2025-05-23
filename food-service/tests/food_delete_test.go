package tests

import (
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
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

	result := db.DB.Create(&testFood)
	assert.NoError(t, result.Error)

	result = db.DB.Delete(&testFood)
	assert.NoError(t, result.Error)

	var foundFood models.Food
	result = db.DB.First(&foundFood, testFood.ID)
	assert.Error(t, result.Error, "Should return error when trying to find deleted record")
}
