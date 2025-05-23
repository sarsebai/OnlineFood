package tests

import (
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryPersistence(t *testing.T) {
	db.DbConnect()

	testCategory := models.Category{
		Name: "Test Category",
	}

	t.Run("Create Category Record", func(t *testing.T) {
		result := db.DB.Create(&testCategory)
		assert.NoError(t, result.Error)
		assert.NotZero(t, testCategory.ID, "ID should not be zero after creation")
	})

	t.Run("Read Category Record", func(t *testing.T) {
		var foundCategory models.Category
		result := db.DB.First(&foundCategory, testCategory.ID)
		assert.NoError(t, result.Error)
		assert.Equal(t, testCategory.Name, foundCategory.Name)
	})

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testCategory)
	})
}
