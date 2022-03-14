package scripts

import (
	"app/utils"
)

var Help = ServerCommand{
	CommandName:  "help",
	Description:  "List all commands and description.",
	Command:      HelpCommand,
	RequiredArgs: []string{},
}

func HelpCommand(args []string) error {
	if len(args) == 0 {
		for name, command := range RegisteredCommands {
			utils.Log().Info("%s - %s", name, command.Description)
		}
	} else {
		if rCommand, ok := RegisteredCommands[args[0]]; ok {
			utils.Log().Info("%s - %s", rCommand.CommandName, rCommand.Description)
		} else {
			utils.Log().Info("Command %s does not exist", rCommand.CommandName)
		}
	}
	return nil
}
