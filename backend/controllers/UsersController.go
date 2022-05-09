package controllers

import (
	"app/helpers/res"
	"app/models/users"
	"strconv"

	"github.com/kataras/iris/v12"
)

type UsersController struct{}

// Create User
// Method:   POST
// Resource: /api/users/create
func (c *UsersController) PostCreate(req users.UserCreateRequest, ctx iris.Context) {
	res.USER_CREATE.Validate(ctx, &req)
	if ctx.IsStopped() {
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

// Edit User
// Method:   POST
// Resource: /api/users/edit
func (c *UsersController) PostEdit(req users.UserEditRequest, ctx iris.Context) {
	res.USER_EDIT.Validate(ctx, &req)
	if ctx.IsStopped() {
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

// Delete User
// Method:   POST
// Resource: /api/users/delete
func (c *UsersController) PostDelete(req users.UserDeleteRequest, ctx iris.Context) {
	res.USER_DELETE.Validate(ctx, &req)
	if ctx.IsStopped() {
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

// Get Users
// Method:   POST
// Resource: /api/users
func (c *UsersController) Get(ctx iris.Context) {
	if users, cannot_list := users.All(); cannot_list != nil {
		res.USERS_LIST.Error(ctx, cannot_list)
	} else {
		res.USERS_LIST.Send(ctx, iris.Map{"users": users})
	}
}
