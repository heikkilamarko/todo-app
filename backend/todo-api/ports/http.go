// Package ports ...
package ports

import (
	"net/http"
	"todo-api/app"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type HTTPServer struct {
	app    *app.App
	router *mux.Router
	logger *zerolog.Logger
}

func NewHTTPServer(app *app.App, router *mux.Router, logger *zerolog.Logger) *HTTPServer {
	s := &HTTPServer{app, router, logger}
	s.registerRoutes()
	return s
}

func (s *HTTPServer) registerRoutes() {
	s.router.HandleFunc("/todos", s.getTodos).
		Methods(http.MethodGet)

	s.router.HandleFunc("/todos", s.createTodo).
		Methods(http.MethodPost)

	s.router.HandleFunc("/todos/{id:[0-9]+}/complete", s.completeTodo).
		Methods(http.MethodPost)
}
