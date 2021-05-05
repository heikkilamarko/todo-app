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
	natsConn   *nats.Conn
	repository *repository
}

// NewController func
func NewController(config *config.Config, logger *zerolog.Logger, db *sql.DB, natsConn *nats.Conn) *Controller {
	repository := newRepository(db, logger)
	return &Controller{config, logger, db, natsConn, repository}
}

// Run method
func (c *Controller) Run() error {

	if err := c.handleTodoCreated(); err != nil {
		return err
	}

	return nil
}
