package controllers

import (
	"app/helpers/res"
	"app/models/users"
	"strconv"

	"github.com/kataras/iris/v12"
)

func CreateUser(ctx iris.Context) {
	type Req struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Redirect string `json:"redirect"`
	}
	var req Req
	if !res.USER_CREATE.Validate(ctx, &req) {
		return
	}

	user := &users.User{
		Username:     req.Username,
		Password:     req.Password,
		Email:        req.Email,
		Locked:       false,
		LastLoggedIn: nil,
	}

	if user_create_err := user.Create(); user_create_err != nil {
		res.USER_CREATE.Error(ctx, user_create_err)
	} else {
		res.USER_CREATE.Send(ctx, iris.Map{"user": user})
	}
}

func EditUser(ctx iris.Context) {
	type Req struct {
		ID       string `json:"id" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Redirect string `json:"redirect"`
	}
	var req Req
	if !res.USER_EDIT.Validate(ctx, &req) {
		return
	}

	id, invalid_id := strconv.ParseUint(req.ID, 10, 64)
	if invalid_id != nil {
		res.USER_EDIT.Error(ctx, invalid_id)
		return
	}

	user := &users.User{
		ID:           id,
		Username:     req.Username,
		Password:     req.Password,
		Email:        req.Email,
		Locked:       false,
		LastLoggedIn: nil,
	}

	if user_update_err := user.Update(); user_update_err != nil {
		res.USER_EDIT.Error(ctx, user_update_err)
	} else {
		res.USER_EDIT.Send(ctx, iris.Map{"user": user})
	}
}

func DeleteUser(ctx iris.Context) {
	type Req struct {
		ID       string `json:"id" validate:"required"`
		Redirect string `json:"redirect"`
	}
	var req Req
	if !res.USER_DELETE.Validate(ctx, &req) {
		return
	}

	user, not_found := users.Find(req.ID)
	if not_found != nil {
		res.USER_DELETE.Error(ctx, not_found)
		return
	}

	if err := user.Delete(); err != nil {
		res.USER_DELETE.Error(ctx, err)
	} else {
		res.USER_DELETE.Send(ctx, iris.Map{})
	}
}

func ListUsers(ctx iris.Context) {
	if users, cannot_list := users.All(); cannot_list != nil {
		res.USER_LIST.Error(ctx, cannot_list)
	} else {
		res.USER_LIST.Send(ctx, iris.Map{"users": users})
	}
}
