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

//go:embed schemas/todo.completed.json
var todoCompletedSchema string

func (c *Controller) handleTodoCompleted(ctx context.Context, m *nats.Msg) {
	m.Ack()

	command := &completeTodoCommand{}

	err := utils.
		NewMessageParser(todoCompletedSchema).
		Parse(m.Data, command)

	if err != nil {
		c.logError(err)

		var message string

		var verr *goutils.ValidationError
		if errors.As(err, &verr) {
			message = verr.Error()
		}

		c.publishError(constants.MessageTodoCompletedError, message)

		return
	}

	if err := c.repository.completeTodo(ctx, command); err != nil {
		c.logError(err)
		c.publishError(constants.MessageTodoCompletedError, "")
		return
	}

	data, err := json.Marshal(command)

	if err != nil {
		c.logError(err)
		return
	}

	if err := c.nc.Publish(constants.MessageTodoCompletedOk, data); err != nil {
		c.logError(err)
		return
	}
}
