package app

import "todo-service/internal/app/command"

type App struct {
	Commands Commands
}

type Commands struct {
	CreateTodo   *command.CreateTodoHandler
	CompleteTodo *command.CompleteTodoHandler
}
