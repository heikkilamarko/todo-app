// Package app provides application level functionality.
package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"os/signal"
	"time"
	"todo-service/app/config"
	"todo-service/app/todos"

	_ "embed"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/xeipuuv/gojsonschema"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

//go:embed schemas/new-todo.json
var newTodoSchema string

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

	schemaLoader := gojsonschema.NewStringLoader(newTodoSchema)

	nc, err := nats.Connect(a.Config.NATSUrl)
	if err != nil {
		a.Logger.Fatal().Err(err).Send()
	}

	nc.Subscribe("todo.created", func(m *nats.Msg) {

		a.Logger.Info().Msgf("Message received (%s)", m.Subject)

		todoLoader := gojsonschema.NewStringLoader(string(m.Data))

		result, err := gojsonschema.Validate(schemaLoader, todoLoader)
		if err != nil {
			a.Logger.Warn().Err(err).Send()
		}

		if !result.Valid() {
			a.Logger.Warn().Msg("Invalid input")
			for _, ve := range result.Errors() {
				a.Logger.Warn().Msgf("Validation error: %s", ve)
			}
			return
		}

		todo := &todos.Todo{}

		if err := json.Unmarshal(m.Data, todo); err != nil {
			a.Logger.Error().Err(err).Send()
			return
		}

		command := &todos.CreateTodoCommand{Todo: todo}

		if err := r.CreateTodo(context.Background(), command); err != nil {
			a.Logger.Error().Err(err).Send()
			return
		}

		todoBytes, err := json.Marshal(todo)

		if err != nil {
			a.Logger.Error().Err(err).Send()
			return
		}

		if err := nc.Publish("todo.processed", todoBytes); err != nil {
			a.Logger.Error().Err(err).Send()
			return
		}
	})

	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)
	<-cs

	err = nc.Drain()
	if err != nil {
		a.Logger.Warn().Err(err).Send()
	}

	nc.Close()

	a.Logger.Info().Msg("Application shutdown")
}
