// Package ports ...
package ports

import (
	"net/http"
	"todo-api/app"

	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
)

type HTTPServer struct {
	app    *app.App
	logger *zerolog.Logger
}

func NewHTTPServer(app *app.App, logger *zerolog.Logger) *HTTPServer {
	return &HTTPServer{app, logger}
}

func (s *HTTPServer) GetTodos(w http.ResponseWriter, r *http.Request) {
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

func (s *HTTPServer) CreateTodo(w http.ResponseWriter, r *http.Request) {
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

func (s *HTTPServer) CompleteTodo(w http.ResponseWriter, r *http.Request) {
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
