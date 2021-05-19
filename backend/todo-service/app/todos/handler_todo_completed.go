package todos

import (
	"context"
	"errors"
	"todo-service/app/utils"

	"github.com/nats-io/nats.go"
)

func (c *Controller) handleTodoCompleted(ctx context.Context, m *nats.Msg) {
	_ = m.Ack()

	command := &completeTodoCommand{}

	if err := c.messageParser.Parse(m, command); err != nil {
		c.logError(err)

		var message string

		var vErr *utils.ValidationError
		if errors.As(err, &vErr) {
			message = vErr.Error()
		}

		_ = c.publishMessage(
			subjectTodoCompletedError,
			errorMessage{
				Code:    subjectTodoCompletedError,
				Message: message,
			},
		)

		return
	}

	if err := c.repository.completeTodo(ctx, command); err != nil {
		c.logError(err)

		_ = c.publishMessage(
			subjectTodoCompletedError,
			errorMessage{
				Code: subjectTodoCompletedError,
			},
		)

		return
	}

	if err := c.publishMessage(subjectTodoCompletedOk, command); err != nil {
		c.logError(err)
	}
}
