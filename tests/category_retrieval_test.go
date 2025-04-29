package tests

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCategories(t *testing.T) {
	db.DbConnect()

	testCategories := []models.Category{
		{
			Name: "Test Category 1",
		},
		{
			Name: "Test Category 2",
		},
		{
			Name: "Test Category 3",
		},
	}

	for i := range testCategories {
		result := db.DB.Create(&testCategories[i])
		assert.NoError(t, result.Error)
		assert.NotZero(t, testCategories[i].ID, "Category ID should not be zero after creation")
	}

	var categories []models.Category
	result := db.DB.Find(&categories)
	assert.NoError(t, result.Error)
	assert.GreaterOrEqual(t, len(categories), len(testCategories), "Should retrieve at least the test categories")

	for _, testCategory := range testCategories {
		found := false
		for _, category := range categories {
			if category.ID == testCategory.ID && category.Name == testCategory.Name {
				found = true
				break
			}
		}
		assert.True(t, found, "Test category should be present in the result")
	}

	t.Cleanup(func() {
		for _, category := range testCategories {
			db.DB.Unscoped().Delete(&category)
		}
	})
}
