package todos

import (
	"context"
	_ "embed"
	"encoding/json"
	"todo-service/app/constants"
	"todo-service/app/utils"
)

//go:embed schemas/todo.completed.json
var todoCompletedSchema string

func (c *Controller) handleTodoCompleted() {

	go func() {

		sub, err := c.js.PullSubscribe(
			constants.MessageTodoCompleted,
			constants.DurableTodoCompleted,
		)

		if err != nil {
			c.logger.Error().Err(err).Send()
			return
		}

		validator := utils.NewJSONSchemaValidator(todoCompletedSchema)

		for {
			ms, err := sub.Fetch(1)
			if err != nil {
				continue
			}

			m := ms[0]

			c.logger.Info().Msgf("message received (%s)", m.Subject)

			ves, err := validator.Validate(string(m.Data))

			if err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			if 0 < len(ves) {
				c.logger.Warn().Msg("invalid message")
				for _, ve := range ves {
					c.logger.Warn().Msgf("validation error: %s", ve)
				}
				continue
			}

			command := &completeTodoCommand{}

			if err := json.Unmarshal(m.Data, command); err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			if err := c.repository.completeTodo(context.Background(), command); err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			data, err := json.Marshal(command)

			if err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			if err := c.nc.Publish(constants.MessageTodoCompletedOk, data); err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			m.Ack()
		}
	}()
}
