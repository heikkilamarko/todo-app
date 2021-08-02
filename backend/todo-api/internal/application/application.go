package application

import (
	"todo-api/internal/application/command"
	"todo-api/internal/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateTodo   *command.CreateTodoHandler
	CompleteTodo *command.CompleteTodoHandler
}

type Queries struct {
	GetTodos *query.GetTodosHandler
}
