package tests

import (
	"app/config"
	"app/scripts"
	"os"

	"github.com/kataras/iris/v12"
)

func CreateApp() *iris.Application {
	config.DotEnv()
	os.Setenv("ENV", "test")
	scripts.RunScripts([]string{"destroy"})
	scripts.RunScripts([]string{"migrate"})
	app := config.Main()
	return app
}

func CreateAppWithAdmin() *iris.Application {
	config.DotEnv()
	os.Setenv("ENV", "test")
	scripts.RunScripts([]string{"destroy"})
	scripts.RunScripts([]string{"migrate"})
	scripts.RunScripts([]string{"create_user", "admin", "password", "test@test.com"})
	app := config.Main()
	return app
}

func DestoryTables() {
	scripts.RunScripts([]string{"destory"})
}
