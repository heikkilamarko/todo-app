import { derived, writable, get } from "svelte/store";
import { toTodo } from "../shared/utils";
import * as api from "../shared/api";
import { showInfo, showError } from "./toasterStore";

/** @type {import("svelte/store").Writable<import("../types").Todo[]>} */
export const todos = writable([]);
export const name = writable("");
export const description = writable("");
export const loading = writable(false);
export const canCreateTodo = derived(name, ($name) => !!$name);

/**
 * @param {number} offset
 * @param {number} limit
 */
export async function getTodos(offset = 0, limit = 10) {
  try {
    loading.set(true);
    const { data } = await api.getTodos({ offset, limit });
    todos.set(data.map(toTodo));
  } catch (e) {
    showError(`todo loading failed\n${e}`);
  } finally {
    loading.set(false);
  }
}

export async function createTodo() {
  try {
    /** @type {import("../types").NewTodo} */
    const todo = {
      name: get(name),
      description: get(description) || null,
    };
    loading.set(true);
    await api.createTodo(todo);
    name.set("");
    description.set("");
    showInfo("todo create job started");
  } catch (e) {
    showError(`todo create job failed\n${e}`);
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
    showInfo("todo complete job started");
  } catch (e) {
    showError(`todo complete job failed\n${e}`);
  } finally {
    loading.set(false);
  }
}
