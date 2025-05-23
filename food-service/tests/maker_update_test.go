package tests

import (
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakerUpdate(t *testing.T) {
	db.DbConnect()

	testMaker := models.Maker{
		Name: "Test Maker",
	}

	result := db.DB.Create(&testMaker)
	assert.NoError(t, result.Error)

	updatedMaker := models.Maker{
		Name: "Updated Maker",
	}

	result = db.DB.Model(&testMaker).Updates(updatedMaker)
	assert.NoError(t, result.Error)

	var foundMaker models.Maker
	result = db.DB.First(&foundMaker, testMaker.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, updatedMaker.Name, foundMaker.Name)

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testMaker)
	})
}
