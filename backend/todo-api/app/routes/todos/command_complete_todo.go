package todos

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/app/constants"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
)

// CompleteTodo method
func (c *Controller) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCompleteTodoRequest(r)

	if err != nil {
		goutils.WriteValidationError(w, err)
		return
	}

	data, err := json.Marshal(command)

	if err != nil {
		c.logger.Error().Err(err).Send()
		return
	}

	_, err = c.js.Publish(constants.MessageTodoCompleted, data)

	if err != nil {
		c.logger.Error().Err(err).Send()
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteResponse(w, http.StatusOK, nil)
}

func parseCompleteTodoRequest(r *http.Request) (*completeTodoCommand, error) {
	validationErrors := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[constants.FieldID])
	if err != nil {
		validationErrors[constants.FieldID] = constants.ErrCodeInvalidTodoID
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &completeTodoCommand{id}, nil
}
