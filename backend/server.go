package main

import (
	"app/config"
	"app/scripts"
	"os"
)

func main() {
	// Run server if there is no other args
	if len(os.Args) == 1 || os.Args[1] == "runserver" {
		app := config.Main()
		app.Listen(":8080")
	} else {
		scripts.RunScripts(os.Args[1:])
	}
}
