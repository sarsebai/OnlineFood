package tests

import (
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryUpdate(t *testing.T) {
	db.DbConnect()

	testCategory := models.Category{
		Name: "Test Category",
	}

	result := db.DB.Create(&testCategory)
	assert.NoError(t, result.Error)

	updatedCategory := models.Category{
		Name: "Updated Category",
	}

	result = db.DB.Model(&testCategory).Updates(updatedCategory)
	assert.NoError(t, result.Error)

	var foundCategory models.Category
	result = db.DB.First(&foundCategory, testCategory.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, updatedCategory.Name, foundCategory.Name)

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testCategory)
	})
}
