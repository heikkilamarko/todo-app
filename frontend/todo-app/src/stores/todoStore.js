import { derived, writable, get } from "svelte/store";
import axios, { AxiosInstance } from "axios";
import { config } from "../shared/config";
import { accessToken } from "../shared/auth";
import { toTodo } from "../shared/utils";
import { showInfo, showError } from "./toasterStore";

/** @type {AxiosInstance} */
let _api;

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
    const r = await api().get(`/todos?offset=${offset}&limit=${limit}`);
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
    await api().post("/todos", todo);
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
    await api().post(`/todos/${id}/complete`);
    showInfo("todo complete job started");
  } catch (e) {
    showError(`todo complete job failed\n${e}`);
  } finally {
    loading.set(false);
  }
}

/**
 * @returns {AxiosInstance}
 */
function api() {
  if (_api) return _api;

  _api = axios.create();

  _api.defaults.baseURL = config.apiUrl;

  _api.interceptors.request.use(
    async (req) => {
      const token = await accessToken();
      req.headers["Authorization"] = `Bearer ${token}`;
      return req;
    },
    (error) => {
      console.log(error);
    }
  );

  return _api;
}
