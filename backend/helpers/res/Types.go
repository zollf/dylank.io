package res

import "github.com/kataras/iris/v12"

type RES_TYPES int64

const (
	undefined RES_TYPES = iota
	ASSETS_LIST
	ASSET_CREATE
	ASSET_DELETE

	TAG_GET
	TAG_CREATE
	TAG_EDIT
	TAG_DELETE
	TAGS_LIST

	PROJECT_GET
	PROJECT_CREATE
	PROJECT_EDIT
	PROJECT_DELETE
	PROJECTS_LIST

	AUTH_LOGIN
	AUTH_LOGOUT

	USERS_LIST
	USER_CREATE
	USER_EDIT
	USER_DELETE
)

type Response struct {
	Success bool           `json:"success"`
	Msg     string         `json:"msg"`
	Path    string         `json:"path"`
	Error   *ResponseError `json:"error"`
	Data    iris.Map       `json:"data"`
}

type ResponseError struct {
	ValidationErrors []ValidationError `json:"validation_errors"`
	Fatal            bool              `json:"fatal"`
	Errors           []string          `json:"errors"`
}
