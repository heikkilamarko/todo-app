import { derived, writable, get } from "svelte/store";
import axios from "axios";
import { showError, showInfo, toTodo } from "../common";
import { config } from "./configStore";

const api = axios.create();

config.subscribe((c) => {
  if (c) api.defaults.baseURL = c.apiUrl;
});

export const todos = writable([]);
export const name = writable("");
export const description = writable("");
export const loading = writable(false);
export const canCreate = derived(name, ($name) => !!$name);

/**
 * @param {number} offset
 * @param {number} limit
 */
export async function load(offset = 0, limit = 10) {
  try {
    loading.set(true);
    var r = await api.get(`/todos?offset=${offset}&limit=${limit}`);
    todos.set((r.data.data ?? []).map(toTodo));
  } catch (e) {
    showError(`todo loading failed\n${e}`);
  } finally {
    loading.set(false);
  }
}

export async function create() {
  try {
    /** @type {import("../types").NewTodo} */
    const todo = {
      name: get(name),
      description: get(description) || null,
    };
    loading.set(true);
    await api.post("/todos", todo);
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
export async function complete(id) {
  try {
    loading.set(true);
    await api.post(`/todos/${id}/complete`);
    showInfo("todo complete job started");
  } catch (e) {
    showError(`todo complete job failed\n${e}`);
  } finally {
    loading.set(false);
  }
}
