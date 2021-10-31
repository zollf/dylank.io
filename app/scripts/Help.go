package scripts

import "log"

var Help = ServerCommand{
	CommandName:  "help",
	Description:  "List all commands and description.",
	Command:      HelpCommand,
	RequiredArgs: []string{},
}

func HelpCommand(args []string) error {
	for name, command := range RegisteredCommands {
		log.Printf("%s - %s", name, command.Description)
	}
	return nil
}
