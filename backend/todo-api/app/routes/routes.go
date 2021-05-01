// Package routes provides routing functionality.
package routes

import (
	"database/sql"
	"net/http"
	"time"
	"todo-api/app/config"
	"todo-api/app/routes/todos"

	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

// RegisterRoutes func
func RegisterRoutes(router *mux.Router, config *config.Config, logger *zerolog.Logger) error {
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

	nec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return err
	}

	pc := todos.NewController(
		todos.NewSQLRepository(db, logger),
		nec,
	)

	router.HandleFunc("/todos", pc.GetTodos).Methods(http.MethodGet)
	router.HandleFunc("/todos", pc.CreateTodo).Methods(http.MethodPost)

	return nil
}
