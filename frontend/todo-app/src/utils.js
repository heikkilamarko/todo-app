import axios from "axios";
import {
  HubConnection,
  HubConnectionBuilder,
  LogLevel,
} from "@microsoft/signalr";

/** @type {import("./types").Config} */
export const config = {
  apiUrl: null,
  notificationMethodName: null,
};

const api = axios.create();

/**
 * Load config from server.
 */
export async function loadConfig() {
  try {
    var response = await api.get("/config.json");
    /** @type {import("./types").Config} */
    var c = response?.data;
    Object.assign(config, c);
    // @ts-ignore
    api.defaults.baseURL = config.apiUrl;
  } catch (e) {
    console.error(e);
    throw e;
  }
}

/**
 * Get todos from server.
 * @returns {Promise<Array<import("./types").Todo>>}
 */
export async function getTodos(offset = 0, limit = 10) {
  try {
    var response = await api.get(`/todos?offset=${offset}&limit=${limit}`);
    /** @type {Array<import("./types").ServerTodo>} */
    var todos = response?.data?.data ?? [];
    return todos.map(toTodo);
  } catch (e) {
    console.error(e);
    throw e;
  }
}

/**
 * Send the new todo to server.
 * @param {import("./types").NewTodo} todo
 */
export async function createTodo(todo) {
  try {
    await api.post("/todos", todo);
  } catch (e) {
    console.error(e);
    throw e;
  }
}

/**
 * Tell server to complete the given todo.
 * @param {number} id
 */
export async function completeTodo(id) {
  try {
    await api.post(`/todos/${id}/complete`);
  } catch (e) {
    console.error(e);
    throw e;
  }
}

/**
 * Get SignalR Hub connection.
 * @returns {HubConnection}
 */
export function getSignalRConnection() {
  return new HubConnectionBuilder()
    .withUrl(`${config.apiUrl}/push/notifications`)
    .configureLogging(LogLevel.Critical)
    .withAutomaticReconnect({
      nextRetryDelayInMilliseconds: () => 5000,
    })
    .build();
}

/**
 * Maps a server todo to a client one.
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
 * Shows an info message.
 * @param {string} message
 */
export function showInfo(message) {
  // @ts-ignore
  window.pushToast(message, "success");
}

/**
 * Shows an error message.
 * @param {string} message
 */
export function showError(message) {
  // @ts-ignore
  window.pushToast(message, "danger");
}
