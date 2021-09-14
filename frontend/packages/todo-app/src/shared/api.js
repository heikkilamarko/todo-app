import ky from "ky";
import { config } from "../shared/config";
import { accessToken } from "../shared/auth";

/**
 * @returns {Promise<string>}
 */
export async function getToken() {
  const { data } = await client().get("todos/token").json();
  return data?.token;
}

/**
 * @param {import("../types").GetTodosRequest} req
 * @returns {Promise<import("../types").GetTodosResponse>}
 */
export async function getTodos(req) {
  return await client()
    .get(`todos?offset=${req.offset}&limit=${req.limit}`)
    .json();
}

/**
 * @param {import("../types").NewTodo} todo
 */
export async function createTodo(todo) {
  await client().post("todos", { json: todo });
}

/**
 * @param {number} id
 */
export async function completeTodo(id) {
  await client().post(`todos/${id}/complete`);
}

/** @type {ky} */
let _client;
function client() {
  return (_client ??= ky.create({
    prefixUrl: config.apiUrl,
    hooks: {
      beforeRequest: [
        async (req) => {
          const token = await accessToken();
          req.headers.set("Authorization", `Bearer ${token}`);
        },
      ],
    },
  }));
}
