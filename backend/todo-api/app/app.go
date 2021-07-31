package app

import (
	"todo-api/app/command"
	"todo-api/app/query"
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
