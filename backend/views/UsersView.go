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
		ID        uint64
		Index     int
		Username  string
		Email     string
		CreatedAt string
		UpdatedAt string
	}

	users, _ := models.GetUsers()
	var users_data []*UserData
	zone, _ := time.LoadLocation("Australia/Perth")

	for i, user := range users {
		users_data = append(users_data, &UserData{
			ID:        user.ID,
			Index:     i + 1,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.In(zone).Format(time.RFC822),
			UpdatedAt: user.UpdatedAt.In(zone).Format(time.RFC822),
		})
	}

	helpers.RenderTemplate(ctx, "users/users", "admin", iris.Map{"Err": err, "Users": users_data})
}

func NewUser(ctx iris.Context) {
	helpers.RenderTemplate(ctx, "users/create", "admin", iris.Map{})
}

func EditUser(ctx iris.Context) {
	type UserData struct {
		ID       uint64
		Username string
		Email    string
	}

	id := ctx.Params().Get("id")
	user, user_err := models.GetUser(id)

	if user_err != nil {
		ctx.View("404")
		return
	}

	user_data := UserData{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	helpers.RenderTemplate(ctx, "users/view", "admin", iris.Map{"User": user_data})
}
