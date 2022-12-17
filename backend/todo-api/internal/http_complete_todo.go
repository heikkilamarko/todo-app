package internal

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type CompleteTodoRequest struct {
	ID int
}

func (req *CompleteTodoRequest) Bind(r *http.Request) error {
	m := make(map[string][]string)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		m["id"] = []string{"invalid id"}
		return ValidationError{m}
	}

	req.ID = id

	return nil
}

type CompleteTodoHandler struct {
	Pub    MessagePublisher
	Logger *zerolog.Logger
}

func (h *CompleteTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &CompleteTodoRequest{}

	if err := req.Bind(r); err != nil {
		h.Logger.Error().Err(err).Send()
		WriteErrorResponse(w, ErrCodeInvalidRequest, err)
		return
	}

	if err := h.Pub.TodoComplete(r.Context(), req.ID); err != nil {
		h.Logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	WriteResponse(w, http.StatusAccepted, nil)
}
