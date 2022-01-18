package tests

import (
	"app/scripts"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func TestValidAuth(t *testing.T) {
	app := CreateApp()
	e := httptest.New(t, app)
	scripts.RunScripts([]string{"create_user", "admin", "password", "test@test.com"})
	e.POST("/api/login").WithForm(iris.Map{
		"username": "admin",
		"password": "password",
	}).Expect().JSON().Object().Value("error").Equal(nil)
	return
}

func TestInvalidAuth(t *testing.T) {
	app := CreateApp()
	e := httptest.New(t, app)
	scripts.RunScripts([]string{"create_user", "admin", "strong_password", "test@test.com"})
	e.POST("/api/login").WithForm(iris.Map{
		"username": "admin",
		"password": "password",
	}).Expect().JSON().Object().Value("error").NotEqual(nil)
	return
}
