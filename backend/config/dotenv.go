package config

import (
	"app/helpers"
	"app/utils"
	"fmt"

	"github.com/joho/godotenv"
)

func DotEnv() {
	root := helpers.RootDir()
	err := godotenv.Load(fmt.Sprintf("%s/.env", root))
	utils.Log().Info("Initialising dotenv")

	if err != nil {
		utils.Log().Error("Failed to initialising dotenv")
	}
}
