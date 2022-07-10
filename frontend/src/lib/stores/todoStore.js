import { writable } from 'svelte/store';
import { toTodo } from '$lib/shared/utils.js';
import * as api from '$lib/shared/api.js';

export default function createStore() {
	const todos = writable([]);

	const loading = writable(false);

	async function getTodos(offset = 0, limit = 10) {
		try {
			loading.set(true);
			const { data } = await api.getTodos({ offset, limit });
			todos.set(data.map(toTodo));
		} finally {
			loading.set(false);
		}
	}

	async function createTodo(todo) {
		try {
			loading.set(true);
			await api.createTodo(todo);
		} finally {
			loading.set(false);
		}
	}

	async function completeTodo(id) {
		try {
			loading.set(true);
			await api.completeTodo(id);
		} finally {
			loading.set(false);
		}
	}

	return { todos, loading, getTodos, createTodo, completeTodo };
}
