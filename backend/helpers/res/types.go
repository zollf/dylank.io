package res

type RES_TYPES int64

const (
	undefined RES_TYPES = iota
	ASSETS_LIST
	ASSETS_CREATE
	ASSETS_DELETE

	TAGS_LIST
	TAGS_CREATE
	TAGS_EDIT
	TAGS_DELETE

	PROJECTS_LIST
	PROJECTS_CREATE
	PROJECTS_EDIT
	PROJECTS_DELETE

	AUTH_LOGIN
	AUTH_LOGOUT
)