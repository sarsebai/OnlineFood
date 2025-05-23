package tests

import (
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakerPersistence(t *testing.T) {
	db.DbConnect()

	testMaker := models.Maker{
		Name: "Test Maker",
	}

	t.Run("Create Maker Record", func(t *testing.T) {
		result := db.DB.Create(&testMaker)
		assert.NoError(t, result.Error)
		assert.NotZero(t, testMaker.ID, "ID should not be zero after creation")
	})

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
