package views

import (
	"app/models"
	"time"

	"github.com/kataras/iris/v12"
)

func Users(ctx iris.Context) {
	err := ctx.URLParam("err")
	success := ctx.URLParam("success")
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

	ctx.View("users/users.pug", iris.Map{"Err": err, "Success": success, "Users": users_data})
}

func NewUser(ctx iris.Context) {
	ctx.View("users/create.pug")
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
		ctx.View("404.pug")
	}

	user_data := UserData{
		ID:          user.ID.Hex(),
		Username:    user.Username,
		Email:       user.Email,
		DateCreated: user.DateCreated,
	}

	ctx.View("users/view.pug", iris.Map{"User": user_data})
}
