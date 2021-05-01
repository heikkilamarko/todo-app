package main

import (
	"os"
	"todo-service/app"
	"todo-service/app/config"

	"github.com/rs/zerolog"
)

func main() {
	c := config.New()
	c.Load()

	zerolog.SetGlobalLevel(c.LogLevel)

	l := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Str("app", c.App).
		Logger()

	l.Info().Str("config", c.String()).Send()

	a := app.New(c, &l)
	a.Run()
}
