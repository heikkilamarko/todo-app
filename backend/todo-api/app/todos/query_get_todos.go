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
		c.logError(err)
		utils.WriteValidationError(w, err)
		return
	}

	todos, err := c.repository.getTodos(r.Context(), query)

	if err != nil {
		c.logError(err)
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteOK(w, todos, query)
}

func parseGetTodosRequest(r *http.Request) (*getTodosQuery, error) {
	errorMap := map[string]string{}

	offset := 0
	limit := utils.LimitMaxPageSize

	var err error = nil

	if value := r.FormValue(utils.FieldPaginationOffset); value != "" {
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			errorMap[utils.FieldPaginationOffset] = utils.ErrCodeInvalidOffset
		}
	}

	if value := r.FormValue(utils.FieldPaginationLimit); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil || limit < 1 || utils.LimitMaxPageSize < limit {
			errorMap[utils.FieldPaginationLimit] = utils.ErrCodeInvalidLimit
		}
	}

	if 0 < len(errorMap) {
		return nil, utils.NewValidationError(errorMap)
	}

	return &getTodosQuery{offset, limit}, nil
}
