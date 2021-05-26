import {
  HubConnection,
  HubConnectionBuilder,
  LogLevel,
} from "@microsoft/signalr";

/**
 * @param {string} url
 * @returns {HubConnection}
 */
export function getSignalRConnection(url) {
  return new HubConnectionBuilder()
    .withUrl(url)
    .configureLogging(LogLevel.Critical)
    .withAutomaticReconnect({
      nextRetryDelayInMilliseconds: () => 5000,
    })
    .build();
}

/**
 * @param {import("./types").ServerTodo} todo
 * @returns {import("./types").Todo}
 */
export function toTodo(todo) {
  return {
    ...todo,
    created_at: new Date(todo.created_at),
    updated_at: new Date(todo.updated_at),
  };
}

/**
 * @param {string} message
 */
export function showInfo(message) {
  // @ts-ignore
  window.pushToast(message, "success");
}

/**
 * @param {string} message
 */
export function showError(message) {
  // @ts-ignore
  window.pushToast(message, "danger");
}
