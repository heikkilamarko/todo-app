package internal

import (
	"net/http"

	"golang.org/x/exp/slog"
)

type GetUserinfoHandler struct {
	Repo   Repository
	Logger *slog.Logger
}

func (h *GetUserinfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ar := GetAuthZResult(r.Context())
	data := Userinfo{ar.Permissions}
	WriteResponse(w, http.StatusOK, NewDataResponse(data, nil))
}
