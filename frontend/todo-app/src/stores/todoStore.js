import { writable } from "svelte/store";
import { toTodo } from "../shared/utils";
import * as api from "../shared/api";

/** @type {import("svelte/store").Writable<import("../types").Todo[]>} */
export const todos = writable([]);
export const loading = writable(false);

/**
 * @param {number} offset
 * @param {number} limit
 */
export async function getTodos(offset = 0, limit = 10) {
  try {
    loading.set(true);
    const { data } = await api.getTodos({ offset, limit });
    todos.set(data.map(toTodo));
  } finally {
    loading.set(false);
  }
}

/**
 * @param {import("../types").NewTodo} todo
 */
export async function createTodo(todo) {
  try {
    loading.set(true);
    await api.createTodo(todo);
  } finally {
    loading.set(false);
  }
}

/**
 * @param {number} id
 */
export async function completeTodo(id) {
  try {
    loading.set(true);
    await api.completeTodo(id);
  } finally {
    loading.set(false);
  }
}
