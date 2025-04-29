package tests

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateCategory(t *testing.T) {
	db.DbConnect()

	testCategory := models.Category{
		Name: "Test Category",
	}
	db.DB.Create(&testCategory)

	newName := "Updated Category"
	result := db.DB.Model(&testCategory).Update("name", newName)
	assert.NoError(t, result.Error)

	var updatedCategory models.Category
	db.DB.First(&updatedCategory, testCategory.ID)
	assert.Equal(t, newName, updatedCategory.Name)

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testCategory)
	})
}
