package migrations

import (
	"app/database"
	"app/models"
	"app/models/assets"
	"app/models/projects"
	"app/models/tags"
)

func Uninstall() {
	db, _ := database.Open()

	db.Migrator().DropTable(&projects.Project{})
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&tags.Tag{})
	db.Migrator().DropTable(&assets.Asset{})
}
