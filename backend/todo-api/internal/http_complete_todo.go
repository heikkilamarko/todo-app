package internal

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	Logger *slog.Logger
}

func (h *CompleteTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &CompleteTodoRequest{}

	if err := req.Bind(r); err != nil {
		h.Logger.Error(err.Error())
		WriteErrorResponse(w, ErrCodeInvalidRequest, err)
		return
	}

	if err := h.Pub.TodoComplete(r.Context(), req.ID); err != nil {
		h.Logger.Error(err.Error())
		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	WriteResponse(w, http.StatusAccepted, nil)
}
