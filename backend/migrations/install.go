package migrations

import (
	"app/database"
	"app/models"
	"app/models/assets"
	"app/models/projects"
	"app/models/tags"
)

func Install() {
	db, _ := database.Open()

	db.AutoMigrate(&projects.Project{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&tags.Tag{})
	db.AutoMigrate(&assets.Asset{})
}
