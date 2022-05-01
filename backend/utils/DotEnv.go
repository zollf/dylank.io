package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func DotEnv() {
	root := RootDir()
	err := godotenv.Load(fmt.Sprintf("%s/.env", root))
	Log().Info("Initialising dotenv")

	if err != nil {
		Log().Error("Failed to initialising dotenv")
	}
}
