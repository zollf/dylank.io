package scripts

import "log"

type ServerCommand struct {
	CommandName  string
	Description  string
	Command      func(args []string) error
	RequiredArgs []string
}

func (command ServerCommand) RunCommand(args []string) {
	if len(args) >= len(command.RequiredArgs) {
		log.Printf("Running command %s", command.CommandName)
		err := command.Command(args)
		if err != nil {
			log.Fatalf("Error running command %s, error: %s", command.CommandName, err)
		} else {
			log.Printf("Finished running command %s", command.CommandName)
		}
	} else {
		log.Fatalf("Command %s requires params: %v", command.CommandName, command.RequiredArgs)
	}
}

// We don't put help as registered command as help uses descriptions from all registered command
var RegisteredCommands = map[string]ServerCommand{
	CreateAdminUser.CommandName: CreateAdminUser,
	Migrate.CommandName:         Migrate,
}

/*
Run command script
MainCommand is always first arg
List of MainCommands
- create_user
*/
func RunScripts(args []string) {
	MainCommand := args[0]

	switch MainCommand {
	case Help.CommandName:
		Help.RunCommand(args[1:])
	default:
		if rCommand, ok := RegisteredCommands[MainCommand]; ok {
			rCommand.RunCommand(args[1:])
		} else {
			log.Println("Please supply a main command.")
		}
	}
}
