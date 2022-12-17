package internal

import (
	"net/http"

	"github.com/rs/zerolog"
)

type GetUserinfoHandler struct {
	AuthZ  AuthZ
	Repo   Repository
	Logger *zerolog.Logger
}

func (h *GetUserinfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ar, err := h.AuthZ.Authorize(r.Context(), &AuthZQuery{
		Token:      GetAccessToken(r.Context()),
		Permission: "todo.read",
	})

	if err != nil {
		h.Logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusUnauthorized, nil)
		return
	}

	if !ar.Allow {
		WriteResponse(w, http.StatusUnauthorized, nil)
		return
	}

	if err != nil {
		h.Logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	data := Userinfo{ar.Permissions}

	WriteResponse(w, http.StatusOK, NewDataResponse(data, nil))
}
