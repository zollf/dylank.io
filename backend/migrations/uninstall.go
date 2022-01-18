package migrations

import (
	"app/database"
	"app/models"
)

func Uninstall() {
	db, _ := database.Open()

	db.Migrator().DropTable(&models.Project{})
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Tag{})
	db.Migrator().DropTable(&models.Asset{})
}
