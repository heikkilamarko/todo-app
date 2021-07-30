package ports

import (
	"net/http"

	"github.com/heikkilamarko/goutils"
)

func (s *HTTPServer) getTodos(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetTodosQuery(r)

	if err != nil {
		s.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	todos, err := s.app.Queries.GetTodos.Handle(r.Context(), query)

	if err != nil {
		s.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, todos, query)
}

func (s *HTTPServer) createTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCreateTodoCommand(r)

	if err != nil {
		s.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := s.app.Commands.CreateTodo.Handle(r.Context(), command); err != nil {
		s.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteResponse(w, http.StatusAccepted, nil)
}

func (s *HTTPServer) completeTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCompleteTodoCommand(r)

	if err != nil {
		s.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := s.app.Commands.CompleteTodo.Handle(r.Context(), command); err != nil {
		s.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteResponse(w, http.StatusAccepted, nil)
}

func (s *HTTPServer) logError(err error) {
	s.logger.Error().Err(err).Send()
}
