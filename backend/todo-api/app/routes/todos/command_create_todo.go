package todos

import (
	"encoding/json"
	"net/http"
	"todo-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// CreateTodo method
func (c *Controller) CreateTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCreateTodoRequest(r)

	if err != nil {
		goutils.WriteValidationError(w, err)
		return
	}

	err = c.natsConn.Publish("todo.created", command.Todo)
	if err != nil {
		goutils.WriteInternalError(w, nil)
	}

	goutils.WriteResponse(w, http.StatusOK, nil)
}

func parseCreateTodoRequest(r *http.Request) (*createTodoCommand, error) {
	validationErrors := map[string]string{}

	todo := &todo{}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		validationErrors[constants.FieldRequestBody] = constants.ErrCodeInvalidPayload
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &createTodoCommand{todo}, nil
}
