// Package todos provides todo functionality.
package todos

import (
	"context"
	"database/sql"
	"encoding/json"
	"todo-service/app/config"
	"todo-service/app/constants"

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
func (c *Controller) Start(ctx context.Context) error {

	sub, err := c.js.PullSubscribe(constants.MessageTodo, constants.DurableTodo)

	if err != nil {
		return err
	}

	go func() {

		c.logger.Info().Msg("todos controller started")

		for {

			select {
			case <-ctx.Done():
				c.logger.Info().Msg("todos controller stopped")
				return
			default:
			}

			ms, err := sub.Fetch(1)
			if err != nil {
				continue
			}

			m := ms[0]

			c.logger.Info().Msgf("message received (%s)", m.Subject)

			switch m.Subject {
			case constants.MessageTodoCreated:
				c.handleTodoCreated(ctx, m)
			case constants.MessageTodoCompleted:
				c.handleTodoCompleted(ctx, m)
			default:
				c.logger.Info().Msgf("unsupported message (%s)", m.Subject)
			}

			c.logger.Info().Msgf("message handled (%s)", m.Subject)
		}
	}()

	return nil
}

func (c *Controller) publishError(subject, message string) {

	data, err := json.Marshal(todoError{subject, message})

	if err != nil {
		c.logError(err)
		return
	}

	if err := c.nc.Publish(subject, data); err != nil {
		c.logError(err)
	}
}

func (c *Controller) logError(err error) {
	c.logger.Error().Err(err).Send()
}
