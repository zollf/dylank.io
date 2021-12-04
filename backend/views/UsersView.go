package views

import (
	"app/helpers"
	"app/models"
	"time"

	"github.com/kataras/iris/v12"
)

func Users(ctx iris.Context) {
	err := ctx.URLParam("err")
	type UserData struct {
		ID          string
		Index       int
		Username    string
		Email       string
		DateCreated string
		DateUpdated string
	}

	users, _ := models.GetUsers()
	var users_data []*UserData
	layout := "2006-01-02 15:04:05 -0700 MST"
	zone, _ := time.LoadLocation("Australia/Perth")

	for i, user := range users {
		DateCreated, _ := time.Parse(layout, user.DateCreated)
		DateUpdated, _ := time.Parse(layout, user.DateUpdated)

		users_data = append(users_data, &UserData{
			ID:          user.ID.Hex(),
			Index:       i + 1,
			Username:    user.Username,
			Email:       user.Email,
			DateCreated: DateCreated.In(zone).Format(time.RFC822),
			DateUpdated: DateUpdated.In(zone).Format(time.RFC822),
		})
	}

	helpers.RenderTemplate(ctx, "users/users", "admin", iris.Map{"Err": err, "Users": users_data})
}

func NewUser(ctx iris.Context) {
	helpers.RenderTemplate(ctx, "users/create", "admin", iris.Map{})
}

func EditUser(ctx iris.Context) {
	type UserData struct {
		ID          string
		Username    string
		Email       string
		DateCreated string
	}

	id := ctx.Params().Get("id")
	user, user_err := models.GetUser(id)

	if user_err != nil {
		ctx.View("404")
		return
	}

	user_data := UserData{
		ID:          user.ID.Hex(),
		Username:    user.Username,
		Email:       user.Email,
		DateCreated: user.DateCreated,
	}

	helpers.RenderTemplate(ctx, "users/view", "admin", iris.Map{"User": user_data})
}
