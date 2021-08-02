package adapters

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/internal/application"
	"todo-api/internal/application/command"
	"todo-api/internal/application/query"
	"todo-api/internal/domain"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
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

type TodoHTTPHandlers struct {
	app    *application.Application
	logger *zerolog.Logger
}

func NewTodoHTTPHandlers(app *application.Application, logger *zerolog.Logger) *TodoHTTPHandlers {
	return &TodoHTTPHandlers{app, logger}
}

// Handlers

func (c *TodoHTTPHandlers) GetTodos(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetTodosQuery(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	todos, err := c.app.Queries.GetTodos.Handle(r.Context(), query)

	if err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, todos, query)
}

func (c *TodoHTTPHandlers) CreateTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCreateTodoCommand(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.app.Commands.CreateTodo.Handle(r.Context(), command); err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteResponse(w, http.StatusAccepted, nil)
}

func (c *TodoHTTPHandlers) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCompleteTodoCommand(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.app.Commands.CompleteTodo.Handle(r.Context(), command); err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteResponse(w, http.StatusAccepted, nil)
}

// Input parsers

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

// Utils

func (c *TodoHTTPHandlers) logError(err error) {
	c.logger.Error().Err(err).Send()
}
