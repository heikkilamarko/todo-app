import axios, { AxiosInstance } from "axios";
import { config } from "../shared/config";
import { accessToken } from "../shared/auth";

/**
 * @param {import("../types").GetTodosRequest} req
 * @returns {Promise<import("../types").GetTodosResponse>}
 */
export async function getTodos(req) {
  const { data } = await client().get(
    `/todos?offset=${req.offset}&limit=${req.limit}`
  );
  return data;
}

/**
 * @param {import("../types").NewTodo} todo
 */
export async function createTodo(todo) {
  await client().post("/todos", todo);
}

/**
 * @param {number} id
 */
export async function completeTodo(id) {
  await client().post(`/todos/${id}/complete`);
}

/** @type {AxiosInstance} */
let _client;

function client() {
  if (_client) return _client;

  _client = axios.create();
  _client.defaults.baseURL = config.apiUrl;
  _client.interceptors.request.use(
    async (req) => {
      const token = await accessToken();
      req.headers["Authorization"] = `Bearer ${token}`;
      return req;
    },
    (error) => {
      console.log(error);
    }
  );

  return _client;
}
