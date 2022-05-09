package utils

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func DotEnv() {
	root := RootDir()
	err := godotenv.Load(fmt.Sprintf("%s/.env", root))
	log.Printf("Initialising dotenv")

	if err != nil {
		log.Printf("Failed to initialising dotenv")
	}
}
