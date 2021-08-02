package application

import "todo-service/internal/application/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	CreateTodo   *command.CreateTodoHandler
	CompleteTodo *command.CompleteTodoHandler
}
