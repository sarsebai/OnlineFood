package tests

import (
	"OnlineFood/db"
	"OnlineFood/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateMaker(t *testing.T) {
	db.DbConnect()

	testMaker := models.Maker{
		Name: "Test Maker",
	}
	db.DB.Create(&testMaker)

	newName := "Updated Maker"
	result := db.DB.Model(&testMaker).Update("name", newName)
	assert.NoError(t, result.Error)

	var updatedMaker models.Maker
	db.DB.First(&updatedMaker, testMaker.ID)
	assert.Equal(t, newName, updatedMaker.Name)

	t.Cleanup(func() {
		db.DB.Unscoped().Delete(&testMaker)
	})
}
