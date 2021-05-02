// Package app provides application level functionality.
package app

import (
	"os"
	"os/signal"
	"todo-service/app/config"
	"todo-service/app/handlers"

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

	if err := handlers.RegisterHandlers(a.Config, a.Logger); err != nil {
		a.Logger.Fatal().Err(err).Send()
	}

	a.Logger.Info().Msg("Application running")

	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)
	<-cs

	a.Logger.Info().Msg("Application shutdown")
}
