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

		c.logInfo("todos controller started")

		for {
			select {
			case <-ctx.Done():
				c.logInfo("todos controller stopped")
				return
			default:
			}

			ms, err := sub.Fetch(1)
			if err != nil {
				continue
			}

			m := ms[0]

			c.logInfo("message received (%s)", m.Subject)

			switch m.Subject {
			case constants.MessageTodoCreated:
				c.handleTodoCreated(ctx, m)
			case constants.MessageTodoCompleted:
				c.handleTodoCompleted(ctx, m)
			default:
				c.logInfo("unsupported message (%s)", m.Subject)
			}

			c.logInfo("message handled (%s)", m.Subject)
		}
	}()

	return nil
}

func (c *Controller) publishMessage(subject string, message interface{}) error {

	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	if err := c.nc.Publish(subject, data); err != nil {
		return err
	}

	return nil
}

func (c *Controller) logInfo(msg string, v ...interface{}) {
	c.logger.Info().Msgf(msg, v...)
}

func (c *Controller) logError(err error) {
	c.logger.Error().Err(err).Send()
}
