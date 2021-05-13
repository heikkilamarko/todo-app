package todos

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"todo-service/app/constants"
	"todo-service/app/utils"

	"github.com/heikkilamarko/goutils"
	"github.com/nats-io/nats.go"
)

//go:embed schemas/todo.created.json
var todoCreatedSchema string

func (c *Controller) handleTodoCreated(ctx context.Context, m *nats.Msg) {
	m.Ack()

	command := &createTodoCommand{}

	err := utils.
		NewMessageParser(todoCreatedSchema).
		Parse(m.Data, command)

	if err != nil {
		c.logError(err)

		var message string

		var verr *goutils.ValidationError
		if errors.As(err, &verr) {
			message = verr.Error()
		}

		c.publishError(constants.MessageTodoCreatedError, message)

		return
	}

	if err := c.repository.createTodo(ctx, command); err != nil {
		c.logError(err)
		c.publishError(constants.MessageTodoCreatedError, "")
		return
	}

	data, err := json.Marshal(command)

	if err != nil {
		c.logError(err)
		return
	}

	if err := c.nc.Publish(constants.MessageTodoCreatedOk, data); err != nil {
		c.logError(err)
		return
	}
}
