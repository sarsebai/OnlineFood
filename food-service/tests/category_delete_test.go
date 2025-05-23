package tests

import (
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryDelete(t *testing.T) {
	db.DbConnect()

	testCategory := models.Category{
		Name: "Test Category",
	}

	result := db.DB.Create(&testCategory)
	assert.NoError(t, result.Error)

	result = db.DB.Delete(&testCategory)
	assert.NoError(t, result.Error)

	var foundCategory models.Category
	result = db.DB.First(&foundCategory, testCategory.ID)
	assert.Error(t, result.Error, "Should return error when trying to find deleted record")
}
