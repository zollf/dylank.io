package config

import (
	"log"

	"github.com/joho/godotenv"
)

func DotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
