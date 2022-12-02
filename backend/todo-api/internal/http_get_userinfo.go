package internal

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type GetUserinfoHandler struct {
	Repo   Repository
	Logger *zerolog.Logger
}

func (h *GetUserinfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	roles := GetRoles(r.Context())

	// TODO: Read role-permission mappings from database

	var permissions []string

	if lo.Contains(roles, "todo-viewer") {
		permissions = []string{"todos.read"}
	}

	if lo.Contains(roles, "todo-user") {
		permissions = []string{"todos.read", "todos.write"}
	}

	data := Userinfo{permissions}

	WriteResponse(w, http.StatusOK, NewDataResponse(data, nil))
}
