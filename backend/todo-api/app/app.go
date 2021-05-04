// Package app provides application level functionality.
package app

import (
	"net/http"
	"todo-api/app/config"
	"todo-api/app/routes"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils/middleware"
	"github.com/ory/graceful"
	"github.com/rs/zerolog"
)

// App struct
type App struct {
	Config *config.Config
	Logger *zerolog.Logger
}

// New func
func New(c *config.Config, l *zerolog.Logger) *App {
	return &App{c, l}
}

// Run method
func (a *App) Run() {
	a.Logger.Info().Msg("application is starting up...")

	router := mux.NewRouter()

	router.Use(
		middleware.Logger(a.Logger),
		middleware.RequestLogger(),
		middleware.ErrorRecovery(),
		middleware.Timeout(a.Config.RequestTimeout),
	)

	if err := routes.RegisterRoutes(router, a.Config, a.Logger); err != nil {
		a.Logger.Fatal().Err(err).Send()
	}

	router.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)

	server := graceful.WithDefaults(&http.Server{
		Addr:    a.Config.Address,
		Handler: router})

	a.Logger.Info().Msgf("application is running at %s", a.Config.Address)

	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		a.Logger.Fatal().Err(err).Send()
	}

	a.Logger.Info().Msg("application is shut down")
}
