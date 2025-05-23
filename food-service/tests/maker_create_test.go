package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
)

func TestCreateMaker(t *testing.T) {
	db.DbConnect()

	testMaker := models.Maker{
		Name: "Test Maker",
	}

	result := db.DB.Create(&testMaker)
	assert.NoError(t, result.Error)
	assert.NotZero(t, testMaker.ID, "ID should not be zero after creation")

	t.Run("Read Maker Record", func(t *testing.T) {
		var foundMaker models.Maker
		result := db.DB.First(&foundMaker, testMaker.ID)
		assert.NoError(t, result.Error)
		assert.Equal(t, testMaker.Name, foundMaker.Name)
	})

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testMaker)
	})
}
