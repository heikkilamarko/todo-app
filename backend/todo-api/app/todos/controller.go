// Package todos provides todo functionality.
package todos

import (
	"database/sql"
	"encoding/json"
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
	return &Controller{config, logger, db, js, &repository{db}}
}

func (c *Controller) publishMessage(subject string, message interface{}) error {
	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	if _, err := c.js.Publish(subject, data); err != nil {
		return err
	}

	return nil
}

func (c *Controller) logError(err error) {
	c.logger.Error().Err(err).Send()
}
