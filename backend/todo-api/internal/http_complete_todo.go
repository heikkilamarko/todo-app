package internal

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type CompleteTodoRequest struct {
	ID int
}

func (req *CompleteTodoRequest) Bind(r *http.Request) error {
	m := make(map[string][]string)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		m["id"] = []string{"invalid id"}
		return ValidationError{m}
	}

	req.ID = id

	return nil
}

type CompleteTodoHandler struct {
	pub    TodoMessagePublisher
	logger *zerolog.Logger
}

func (h *CompleteTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := AuthorizeWrite(r); err != nil {
		h.logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusUnauthorized, nil)
		return
	}

	req := &CompleteTodoRequest{}

	if err := req.Bind(r); err != nil {
		h.logger.Error().Err(err).Send()
		WriteErrorResponse(w, ErrCodeInvalidRequest, err)
		return
	}

	if err := h.pub.TodoComplete(r.Context(), req.ID); err != nil {
		h.logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	WriteResponse(w, http.StatusAccepted, nil)
}
