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
	roles := GetRoles(r.Context())

	permissions, err := h.Repo.GetPermissions(r.Context(), roles)

	if err != nil {
		h.Logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	data := Userinfo{permissions}

	WriteResponse(w, http.StatusOK, NewDataResponse(data, nil))
}
