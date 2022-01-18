package scripts

import (
	"app/migrations"
)

var Destroy = ServerCommand{
	CommandName:  "destroy",
	Description:  "Destroy all tables",
	Command:      DestroyCommand,
	RequiredArgs: []string{},
}

func DestroyCommand(args []string) error {
	migrations.Uninstall()
	return nil
}
