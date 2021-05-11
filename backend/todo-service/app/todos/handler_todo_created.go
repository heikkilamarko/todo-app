package todos

import (
	"context"
	_ "embed"
	"encoding/json"
	"todo-service/app/constants"
	"todo-service/app/utils"
)

//go:embed schemas/todo.created.json
var todoCreatedSchema string

func (c *Controller) handleTodoCreated() {

	go func() {

		sub, err := c.js.PullSubscribe(
			constants.MessageTodoCreated,
			constants.DurableTodoCreated,
		)

		if err != nil {
			c.logger.Error().Err(err).Send()
			return
		}

		validator := utils.NewJSONSchemaValidator(todoCreatedSchema)

		for {
			ms, err := sub.Fetch(1)
			if err != nil {
				continue
			}

			m := ms[0]
			m.Ack()

			c.logger.Info().Msgf("message received (%s)", m.Subject)

			ves, err := validator.Validate(string(m.Data))

			if err != nil {
				c.publishTodoCreatedError(err, "")
				continue
			}

			if 0 < len(ves) {
				c.logger.Warn().Msg("invalid message")
				for _, ve := range ves {
					c.logger.Warn().Msgf("validation error: %s", ve)
				}
				c.publishTodoCreatedError(nil, "bad request")
				continue
			}

			todo := &todo{}

			if err := json.Unmarshal(m.Data, todo); err != nil {
				c.publishTodoCreatedError(err, "")
				continue
			}

			command := &createTodoCommand{Todo: todo}

			if err := c.repository.createTodo(context.Background(), command); err != nil {
				c.publishTodoCreatedError(err, "")
				continue
			}

			data, err := json.Marshal(todo)

			if err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			if err := c.nc.Publish(constants.MessageTodoCreatedOk, data); err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}
		}
	}()
}

func (c *Controller) publishTodoCreatedError(err error, message string) {

	if err != nil {
		c.logger.Error().Err(err).Send()
	}

	data, err := json.Marshal(todoError{constants.MessageTodoCreatedError, message})

	if err != nil {
		c.logger.Error().Err(err).Send()
		return
	}

	if err := c.nc.Publish(constants.MessageTodoCreatedError, data); err != nil {
		c.logger.Error().Err(err).Send()
	}
}
