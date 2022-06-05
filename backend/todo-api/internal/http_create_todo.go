package internal

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

type CreateTodoRequest struct {
	Todo *Todo
}

func (req *CreateTodoRequest) Bind(r *http.Request) error {
	todo := &Todo{}

	m := make(map[string][]string)

	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		m["request_body"] = []string{"invalid request payload"}
		return ValidationError{m}
	}

	req.Todo = todo

	return nil
}

type CreateTodoHandler struct {
	Pub    MessagePublisher
	Logger *zerolog.Logger
}

func (h *CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := AuthorizeWrite(r); err != nil {
		h.Logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusUnauthorized, nil)
		return
	}

	req := &CreateTodoRequest{}

	if err := req.Bind(r); err != nil {
		h.Logger.Error().Err(err).Send()
		WriteErrorResponse(w, ErrCodeInvalidRequest, err)
		return
	}

	if err := h.Pub.TodoCreate(r.Context(), req.Todo); err != nil {
		h.Logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	WriteResponse(w, http.StatusAccepted, nil)
}
