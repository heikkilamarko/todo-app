package todos

import (
	"net/http"
	"strconv"
	"todo-api/app/utils"
)

// GetTodos method
func (c *Controller) GetTodos(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetTodosRequest(r)

	if err != nil {
		utils.WriteValidationError(w, err)
		return
	}

	todos, err := c.repository.getTodos(r.Context(), query)

	if err != nil {
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteOK(w, todos, query)
}

func parseGetTodosRequest(r *http.Request) (*getTodosQuery, error) {
	validationErrors := map[string]string{}

	var offset int = 0
	var limit int = utils.LimitMaxPageSize

	var err error = nil

	if value := r.FormValue(utils.FieldPaginationOffset); value != "" {
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			validationErrors[utils.FieldPaginationOffset] = utils.ErrCodeInvalidOffset
		}
	}

	if value := r.FormValue(utils.FieldPaginationLimit); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil || limit < 1 || utils.LimitMaxPageSize < limit {
			validationErrors[utils.FieldPaginationLimit] = utils.ErrCodeInvalidLimit
		}
	}

	if 0 < len(validationErrors) {
		return nil, utils.NewValidationError(validationErrors)
	}

	return &getTodosQuery{offset, limit}, nil
}
