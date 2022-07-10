import { writable } from 'svelte/store';
import Centrifuge from 'centrifuge';
import * as api from '$lib/shared/api.js';
import { stores } from '$lib/shared/stores.js';

export default function createStore() {
	const {
		config,
		toasterStore: { showError },
		todoStore: { getTodos }
	} = stores;

	const connected = writable(null);

	async function connect() {
		let token;

		try {
			token = await api.getToken();
		} catch (error) {
			showError(`real-time connection error\n${error}`);
			return;
		}

		const centrifuge = new Centrifuge(config.notificationsUrl);
		centrifuge.setToken(token);

		centrifuge.on('connect', () => connected.set(true));
		centrifuge.on('disconnect', () => connected.set(false));

		centrifuge.subscribe('notifications', async (ctx) => {
			const { type, data } = ctx.data ?? {};
			switch (type) {
				case 'todo.create.ok':
				case 'todo.complete.ok':
					try {
						await getTodos();
					} catch (error) {
						showError(`todo loading failed\n${error}`);
					}
					break;
				case 'todo.create.error':
				case 'todo.complete.error':
					showError(`error: ${data.code}\n${data.message || '-'}`);
					break;
			}
		});

		try {
			centrifuge.connect();
			return () => centrifuge.disconnect();
		} catch (error) {
			connected.set(false);
			showError(`real-time connection error\n${error}`);
		}
	}

	return { connected, connect };
}
