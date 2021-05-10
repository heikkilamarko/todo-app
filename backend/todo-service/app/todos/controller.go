// Package todos provides todo functionality.
package todos

import (
	"database/sql"
	"todo-service/app/config"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

// Controller struct
type Controller struct {
	config     *config.Config
	logger     *zerolog.Logger
	db         *sql.DB
	nc         *nats.Conn
	js         nats.JetStreamContext
	repository *repository
}

// NewController func
func NewController(config *config.Config, logger *zerolog.Logger, db *sql.DB, nc *nats.Conn, js nats.JetStreamContext) *Controller {
	repository := newRepository(db, logger)
	return &Controller{config, logger, db, nc, js, repository}
}

// Start method
func (c *Controller) Start() error {
	c.handleTodoCreated()
	c.handleTodoCompleted()
	return nil
}
