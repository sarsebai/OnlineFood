package tests

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteCategory(t *testing.T) {
	db.DbConnect()

	testCategory := models.Category{
		Name: "Test Category",
	}
	db.DB.Create(&testCategory)

	result := db.DB.Delete(&testCategory)
	assert.NoError(t, result.Error)

	var deletedCategory models.Category
	result = db.DB.First(&deletedCategory, testCategory.ID)
	assert.Error(t, result.Error, "Record should not be found after deletion")

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testCategory)
	})
}
