package todos

import (
	"context"
	"errors"
	"todo-service/app/utils"

	"github.com/heikkilamarko/goutils"
	"github.com/nats-io/nats.go"
)

func (c *Controller) handleTodoCreated(ctx context.Context, m *nats.Msg) {
	m.Ack()

	command := &createTodoCommand{}

	err := utils.
		NewMessageParser(schemaTodoCreated).
		Parse(m.Data, command)

	if err != nil {
		c.logError(err)

		var message string

		var verr *goutils.ValidationError
		if errors.As(err, &verr) {
			message = verr.Error()
		}

		c.publishMessage(
			subjectTodoCreatedError,
			errorMessage{subjectTodoCreatedError, message},
		)

		return
	}

	if err := c.repository.createTodo(ctx, command); err != nil {
		c.logError(err)
		c.publishMessage(
			subjectTodoCreatedError,
			errorMessage{Code: subjectTodoCreatedError},
		)
		return
	}

	if err := c.publishMessage(subjectTodoCreatedOk, command); err != nil {
		c.logError(err)
	}
}
