package tests

import (
	"OnlineFood/food-service/db"
	"OnlineFood/food-service/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakerDelete(t *testing.T) {
	db.DbConnect()

	testMaker := models.Maker{
		Name: "Test Maker",
	}

	result := db.DB.Create(&testMaker)
	assert.NoError(t, result.Error)

	result = db.DB.Delete(&testMaker)
	assert.NoError(t, result.Error)

	var foundMaker models.Maker
	result = db.DB.First(&foundMaker, testMaker.ID)
	assert.Error(t, result.Error, "Should return error when trying to find deleted record")
}
