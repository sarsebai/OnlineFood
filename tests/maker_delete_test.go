package tests

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteMaker(t *testing.T) {
	db.DbConnect()

	testMaker := models.Maker{
		Name: "Test Maker",
	}
	db.DB.Create(&testMaker)

	result := db.DB.Delete(&testMaker)
	assert.NoError(t, result.Error)

	var deletedMaker models.Maker
	result = db.DB.First(&deletedMaker, testMaker.ID)
	assert.Error(t, result.Error, "Record should not be found after deletion")

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testMaker)
	})
}
