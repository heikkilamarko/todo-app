export const Notification = {
  TodoCreatedOk: "todo.created.ok",
  TodoCreatedError: "todo.created.error",
  TodoCompletedOk: "todo.completed.ok",
  TodoCompletedError: "todo.completed.error",
};

/**
 * @param {import("../types").ServerTodo} todo
 * @returns {import("../types").Todo}
 */
export function toTodo(todo) {
  return {
    ...todo,
    created_at: new Date(todo.created_at),
    updated_at: new Date(todo.updated_at),
  };
}
