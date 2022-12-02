import ky from 'ky';
import { stores } from './stores.js';

export function getUserinfo() {
	return todoApi().get('todos/userinfo').json();
}

export async function getToken() {
	const { data } = await todoApi().get('todos/token').json();
	return data?.token;
}

export async function getTodos(req) {
	return await todoApi().get('todos', { searchParams: req }).json();
}

export async function createTodo(todo) {
	await todoApi().post('todos', { json: todo });
}

export async function completeTodo(id) {
	await todoApi().post(`todos/${id}/complete`);
}

function todoApi() {
	return ky.create({
		prefixUrl: stores.config.apiUrl,
		retry: {
			limit: 0
		},
		hooks: {
			beforeRequest: [setBearerToken]
		}
	});
}

async function setBearerToken(req) {
	if (!stores.auth.isSignedIn()) return;
	const token = await stores.auth.getAccessToken();
	req.headers.set('Authorization', `Bearer ${token}`);
}
