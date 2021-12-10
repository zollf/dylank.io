package scripts

import (
	"app/migrations"
)

// Used to create admin user on startup
// Useful for dev
var Migrate = ServerCommand{
	CommandName:  "migrate",
	Description:  "Migrates database with all models",
	Command:      MigrateCommand,
	RequiredArgs: []string{},
}

func MigrateCommand(args []string) error {
	migrations.Install()
	return nil
}
