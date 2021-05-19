package todos

import (
	"context"
	"errors"
	"todo-service/app/utils"

	"github.com/nats-io/nats.go"
)

func (c *Controller) handleTodoCreated(ctx context.Context, m *nats.Msg) {
	_ = m.Ack()

	command := &createTodoCommand{}

	if err := c.messageParser.Parse(m, command); err != nil {
		c.logError(err)

		var message string

		var vErr *utils.ValidationError
		if errors.As(err, &vErr) {
			message = vErr.Error()
		}

		_ = c.publishMessage(
			subjectTodoCreatedError,
			errorMessage{
				Code:    subjectTodoCreatedError,
				Message: message,
			},
		)

		return
	}

	if err := c.repository.createTodo(ctx, command); err != nil {
		c.logError(err)

		_ = c.publishMessage(
			subjectTodoCreatedError,
			errorMessage{
				Code: subjectTodoCreatedError,
			},
		)

		return
	}

	if err := c.publishMessage(subjectTodoCreatedOk, command); err != nil {
		c.logError(err)
	}
}
