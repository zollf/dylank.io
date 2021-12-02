package scripts

import "log"

var Help = ServerCommand{
	CommandName:  "help",
	Description:  "List all commands and description.",
	Command:      HelpCommand,
	RequiredArgs: []string{},
}

func HelpCommand(args []string) error {
	if len(args) == 0 {
		for name, command := range RegisteredCommands {
			log.Printf("%s - %s", name, command.Description)
		}
	} else {
		if rCommand, ok := RegisteredCommands[args[0]]; ok {
			log.Printf("%s - %s", rCommand.CommandName, rCommand.Description)
		} else {
			log.Printf("Command %s does not exist", rCommand.CommandName)
		}
	}
	return nil
}
