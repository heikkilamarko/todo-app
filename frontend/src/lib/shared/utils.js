export function toTodo(todo) {
	return {
		...todo,
		created_at: new Date(todo.created_at),
		updated_at: new Date(todo.updated_at)
	};
}
