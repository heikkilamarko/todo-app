package todos

import (
	"context"
	_ "embed"
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

		c.publishMessage(
			constants.MessageTodoCreatedError,
			errorMessage{constants.MessageTodoCreatedError, message},
		)

		return
	}

	if err := c.repository.createTodo(ctx, command); err != nil {
		c.logError(err)
		c.publishMessage(
			constants.MessageTodoCreatedError,
			errorMessage{Code: constants.MessageTodoCreatedError},
		)
		return
	}

	if err := c.publishMessage(constants.MessageTodoCreatedOk, command); err != nil {
		c.logError(err)
	}
}
