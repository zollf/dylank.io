package scripts

import (
	"app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Used to create admin user on startup
// Useful for dev
var CreateAdminUser = ServerCommand{
	CommandName:  "create_user",
	Description:  "Creates user. Usage: create_user <username> <password> <email>",
	Command:      CreateAdminUserCommand,
	RequiredArgs: []string{"username", "password", "email"},
}

func CreateAdminUserCommand(args []string) error {
	username := args[0]
	password := args[1]
	email := args[2]

	err := models.CreateUser(&models.User{
		ID:           primitive.NewObjectID(),
		Username:     username,
		Password:     password,
		Email:        email,
		Locked:       false,
		DateCreated:  time.Now().UTC().String(),
		DateUpdated:  time.Now().UTC().String(),
		LastLoggedIn: "",
	})

	return err
}
