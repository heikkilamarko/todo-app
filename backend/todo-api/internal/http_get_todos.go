package internal

import (
	"net/http"
	"strconv"

	"github.com/rs/zerolog"
)

type GetTodosRequest struct {
	Query *GetTodosQuery
}

func (req *GetTodosRequest) Bind(r *http.Request) error {
	m := make(map[string][]string)

	offset := 0
	limit := 100

	var err error

	if value := r.FormValue("offset"); value != "" {
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			m["offsett"] = []string{"invalid offset"}
		}
	}

	if value := r.FormValue("limit"); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil || limit < 1 || 100 < limit {
			m["limit"] = []string{"invalid limit"}
		}
	}

	if 0 < len(m) {
		return ValidationError{m}
	}

	req.Query = &GetTodosQuery{offset, limit}

	return nil
}

type GetTodosHandler struct {
	Repo   Repository
	Logger *zerolog.Logger
}

func (h *GetTodosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := AuthorizeRead(r); err != nil {
		h.Logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusUnauthorized, nil)
		return
	}

	req := &GetTodosRequest{}

	if err := req.Bind(r); err != nil {
		h.Logger.Error().Err(err).Send()
		WriteErrorResponse(w, ErrCodeInvalidRequest, err)
		return
	}

	data, err := h.Repo.GetTodos(r.Context(), req.Query)

	if err != nil {
		h.Logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	WriteResponse(w, http.StatusOK, NewDataResponse(data, req.Query))
}
