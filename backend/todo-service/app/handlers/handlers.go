// Package handlers provides message handling functionality.
package handlers

import (
	"database/sql"
	"time"
	"todo-service/app/config"
	"todo-service/app/handlers/todos"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

func RegisterHandlers(config *config.Config, logger *zerolog.Logger) error {
	db, err := sql.Open("pgx", config.DBConnectionString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return err
	}

	nc, err := nats.Connect(config.NATSUrl)
	if err != nil {
		return err
	}

	tc := todos.NewController(
		todos.NewSQLRepository(db, logger),
		nc,
		logger,
	)

	if err := tc.HandleTodoCreated(); err != nil {
		logger.Fatal().Err(err).Send()
	}

	return nil
}
