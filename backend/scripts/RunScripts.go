package scripts

import (
	"app/utils"
	"os"
)

type ServerCommand struct {
	CommandName  string
	Description  string
	Command      func(args []string) error
	RequiredArgs []string
}

func (command ServerCommand) RunCommand(args []string) {
	if len(args) >= len(command.RequiredArgs) {
		if os.Getenv("ENV") != "test" {
			utils.Log().Info("Running command %s", command.CommandName)
		}
		err := command.Command(args)
		if err != nil {
			utils.Log().Error("Error running command %s, error: %s", command.CommandName, err)
		} else {
			if os.Getenv("ENV") != "test" {
				utils.Log().Info("Finished running command %s", command.CommandName)
			}
		}
	} else {
		utils.Log().Error("Command %s requires params: %v", command.CommandName, command.RequiredArgs)
	}
}

// We don't put help as registered command as help uses descriptions from all registered command
var RegisteredCommands = map[string]ServerCommand{
	CreateAdminUser.CommandName: CreateAdminUser,
	Migrate.CommandName:         Migrate,
	Destroy.CommandName:         Destroy,
}

// Run command script, main command is always first arg
func RunScripts(args []string) {
	utils.DotEnv()

	MainCommand := args[0]

	switch MainCommand {
	case Help.CommandName:
		Help.RunCommand(args[1:])
	default:
		if rCommand, ok := RegisteredCommands[MainCommand]; ok {
			rCommand.RunCommand(args[1:])
		} else {
			if os.Getenv("ENV") != "test" {
				utils.Log().Info("Please supply a main command.")
			}
		}
	}
}
