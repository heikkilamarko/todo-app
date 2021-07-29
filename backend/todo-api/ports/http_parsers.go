package ports

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/app/command"
	"todo-api/app/query"
	"todo-api/domain"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
)

const (
	errCodeInvalidID          = "invalid_id"
	errCodeInvalidOffset      = "invalid_offset"
	errCodeInvalidLimit       = "invalid_limit"
	errCodeInvalidRequestBody = "invalid_request_body"
)

const (
	fieldID               = "id"
	fieldPaginationOffset = "offset"
	fieldPaginationLimit  = "limit"
	fieldRequestBody      = "request_body"
)

const (
	limitMaxPageSize = 100
)

func parseGetTodosQuery(r *http.Request) (*query.GetTodos, error) {
	errorMap := map[string]string{}

	offset := 0
	limit := limitMaxPageSize

	var err error

	if value := r.FormValue(fieldPaginationOffset); value != "" {
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			errorMap[fieldPaginationOffset] = errCodeInvalidOffset
		}
	}

	if value := r.FormValue(fieldPaginationLimit); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil || limit < 1 || limitMaxPageSize < limit {
			errorMap[fieldPaginationLimit] = errCodeInvalidLimit
		}
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &query.GetTodos{
		Offset: offset,
		Limit:  limit,
	}, nil
}

func parseCreateTodoCommand(r *http.Request) (*command.CreateTodo, error) {
	errorMap := map[string]string{}

	todo := &domain.Todo{}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		errorMap[fieldRequestBody] = errCodeInvalidRequestBody
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &command.CreateTodo{
		Todo: todo,
	}, nil
}

func parseCompleteTodoCommand(r *http.Request) (*command.CompleteTodo, error) {
	errorMap := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[fieldID])
	if err != nil {
		errorMap[fieldID] = errCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &command.CompleteTodo{
		ID: id,
	}, nil
}
