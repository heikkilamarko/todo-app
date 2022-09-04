import ky from 'ky';
import { getAccessToken, isSignedIn } from './auth.js';
import { stores } from './stores.js';

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
		hooks: {
			beforeRequest: [setBearerToken]
		}
	});
}

async function setBearerToken(req) {
	if (!isSignedIn()) return;
	const token = await getAccessToken();
	req.headers.set('Authorization', `Bearer ${token}`);
}
