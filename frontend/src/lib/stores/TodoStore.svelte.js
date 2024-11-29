import { toTodo } from '$lib/shared/utils.js';
import * as api from '$lib/shared/api.js';

export class TodoStore {
	todos = $state([]);
	loading = $state(false);

	async getTodos(offset = 0, limit = 10) {
		try {
			this.loading = true;
			const { data } = await api.getTodos({ offset, limit });
			this.todos = data.map(toTodo);
		} finally {
			this.loading = false;
		}
	}

	async createTodo(todo) {
		try {
			this.loading = true;
			await api.createTodo(todo);
		} finally {
			this.loading = false;
		}
	}

	async completeTodo(id) {
		try {
			this.loading = true;
			await api.completeTodo(id);
		} finally {
			this.loading = false;
		}
	}
}
