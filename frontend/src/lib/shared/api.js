import ky from 'ky';
import { accessToken } from './auth.js';
import { stores } from './stores.js';

export async function getToken() {
	const { data } = await client().get('todos/token').json();
	return data?.token;
}

export async function getTodos(req) {
	return await client().get('todos', { searchParams: req }).json();
}

export async function createTodo(todo) {
	await client().post('todos', { json: todo });
}

export async function completeTodo(id) {
	await client().post(`todos/${id}/complete`);
}

let _client;

function client() {
	return (_client ??= ky.create({
		prefixUrl: stores.config.apiUrl,
		hooks: {
			beforeRequest: [
				async (req) => {
					const token = await accessToken();
					req.headers.set('Authorization', `Bearer ${token}`);
				}
			]
		}
	}));
}
