package migrations

import (
	"app/database"
	"app/models"
)

func Install() {
	db, _ := database.Open()

	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.Asset{})
}
