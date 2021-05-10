// Package todos provides todo functionality.
package todos

import (
	"database/sql"
	"todo-api/app/config"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

// Controller struct
type Controller struct {
	config     *config.Config
	logger     *zerolog.Logger
	db         *sql.DB
	js         nats.JetStreamContext
	repository *repository
}

// NewController func
func NewController(config *config.Config, logger *zerolog.Logger, db *sql.DB, js nats.JetStreamContext) *Controller {
	repository := newRepository(db, logger)
	return &Controller{config, logger, db, js, repository}
}
