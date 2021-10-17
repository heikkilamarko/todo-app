/**
 * @param {import("../../types").ServerTodo} todo server todo
 * @returns {import("../../types").Todo} todo
 */
export function toTodo(todo) {
  return {
    ...todo,
    created_at: new Date(todo.created_at),
    updated_at: new Date(todo.updated_at),
  };
}
