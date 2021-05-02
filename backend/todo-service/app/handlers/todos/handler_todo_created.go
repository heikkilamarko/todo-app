package todos

import (
	"context"
	_ "embed"
	"encoding/json"
	"todo-service/app/handlers/todos/schemas"

	"github.com/nats-io/nats.go"
)

const subjectTodoCreated = "todo.created"
const subjectTodoProcessed = "todo.processed"

//go:embed schemas/new-todo.json
var newTodoSchema string

func (c *Controller) HandleTodoCreated() error {

	validator := schemas.NewValidator(newTodoSchema)

	_, err := c.NatsConn.Subscribe(subjectTodoCreated, func(m *nats.Msg) {

		c.Logger.Info().Msgf("Message received (%s)", m.Subject)

		ves, err := validator.Validate(string(m.Data))

		if err != nil {
			c.Logger.Error().Err(err).Send()
			return
		}

		if 0 < len(ves) {
			c.Logger.Warn().Msg("Invalid message")
			for _, ve := range ves {
				c.Logger.Warn().Msgf("Validation error: %s", ve)
			}
			return
		}

		todo := &Todo{}

		if err := json.Unmarshal(m.Data, todo); err != nil {
			c.Logger.Error().Err(err).Send()
			return
		}

		command := &CreateTodoCommand{Todo: todo}

		if err := c.Repository.CreateTodo(context.Background(), command); err != nil {
			c.Logger.Error().Err(err).Send()
			return
		}

		todoBytes, err := json.Marshal(todo)

		if err != nil {
			c.Logger.Error().Err(err).Send()
			return
		}

		if err := c.NatsConn.Publish(subjectTodoProcessed, todoBytes); err != nil {
			c.Logger.Error().Err(err).Send()
			return
		}
	})

	if err != nil {
		return err
	}

	return nil
}
