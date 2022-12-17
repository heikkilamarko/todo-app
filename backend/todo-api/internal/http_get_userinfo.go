package internal

import (
	"net/http"

	"github.com/rs/zerolog"
)

type GetUserinfoHandler struct {
	Repo   Repository
	Logger *zerolog.Logger
}

func (h *GetUserinfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ar := GetAuthZResult(r.Context())
	data := Userinfo{ar.Permissions}
	WriteResponse(w, http.StatusOK, NewDataResponse(data, nil))
}
