// Package app provides application level functionality.
package app

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"time"
	"todo-service/app/config"
	"todo-service/app/todos"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
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
	a.Logger.Info().Msgf("Application running")

	// Init database

	db, err := sql.Open("pgx", a.Config.DBConnectionString)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	r := todos.NewSQLRepository(db, a.Logger)

	// Message handling

	nc, err := nats.Connect(a.Config.NATSUrl)
	if err != nil {
		a.Logger.Fatal().Err(err).Send()
	}

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	c.Subscribe("todo.created", func(todo *todos.Todo) {
		a.Logger.Info().Msgf("Received a person: %+v\n", todo)

		command := &todos.CreateTodoCommand{Todo: todo}

		if err := r.CreateTodo(context.Background(), command); err != nil {
			a.Logger.Error().Err(err).Send()
			return
		}

		if err := c.Publish("todo.processed", todo); err != nil {
			a.Logger.Error().Err(err).Send()
			return
		}
	})

	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)
	<-cs

	a.Logger.Info().Msg("Draining...")

	err = nc.Drain()
	if err != nil {
		a.Logger.Warn().Err(err).Send()
	}

	a.Logger.Info().Msg("Closing...")

	nc.Close()

	a.Logger.Info().Msg("Application shutdown")
}
