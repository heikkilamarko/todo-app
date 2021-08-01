package app

import (
	"todo-api/internal/app/command"
	"todo-api/internal/app/query"
)

type App struct {
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
