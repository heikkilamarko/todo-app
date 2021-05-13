package todos

import (
	"context"
	_ "embed"
	"encoding/json"
	"todo-service/app/constants"
	"todo-service/app/utils"

	"github.com/nats-io/nats.go"
)

//go:embed schemas/todo.completed.json
var todoCompletedSchema string

func (c *Controller) handleTodoCompleted(ctx context.Context, m *nats.Msg) {
	m.Ack()

	validator := utils.NewJSONSchemaValidator(todoCompletedSchema)

	ves, err := validator.Validate(string(m.Data))

	if err != nil {
		c.logError(err)
		c.publishError(constants.MessageTodoCompletedError, "")
		return
	}

	if 0 < len(ves) {
		c.logInfo("invalid message")
		for _, ve := range ves {
			c.logInfo("validation error: %s", ve)
		}
		c.publishError(constants.MessageTodoCompletedError, "bad request")
		return
	}

	command := &completeTodoCommand{}

	if err := json.Unmarshal(m.Data, command); err != nil {
		c.logError(err)
		c.publishError(constants.MessageTodoCompletedError, "")
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
