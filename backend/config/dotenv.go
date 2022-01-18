package config

import (
	"app/helpers"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func DotEnv() {
	root := helpers.RootDir()
	err := godotenv.Load(fmt.Sprintf("%s/.env", root))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
