package internal

import (
	"encoding/json"
	"net/http"

	"golang.org/x/exp/slog"
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
	Logger *slog.Logger
}

func (h *CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &CreateTodoRequest{}

	if err := req.Bind(r); err != nil {
		h.Logger.Error(err.Error())
		WriteErrorResponse(w, ErrCodeInvalidRequest, err)
		return
	}

	if err := h.Pub.TodoCreate(r.Context(), req.Todo); err != nil {
		h.Logger.Error(err.Error())
		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	WriteResponse(w, http.StatusAccepted, nil)
}
