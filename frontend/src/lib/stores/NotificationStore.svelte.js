import { Centrifuge } from 'centrifuge';
import * as api from '$lib/shared/api.js';
import { stores } from '$lib/shared/stores.js';

export class NotificationStore {
	connected = $state(null);

	async connect() {
		let token;
		let centrifuge;

		const { config, toasterStore, todoStore } = stores;

		try {
			token = await api.getToken();
		} catch (error) {
			toasterStore.showError(`real-time connection error\n${error}`);
			return;
		}

		this.centrifuge = new Centrifuge(config.notificationsUrl, { token });

		const sub = this.centrifuge.newSubscription('notifications');

		sub.on('publication', async (ctx) => {
			const { type, data } = ctx.data ?? {};
			switch (type) {
				case 'todo.create.ok':
				case 'todo.complete.ok':
					try {
						await todoStore.getTodos();
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
		sub.on('subscribed', () => (this.connected = true));
		sub.on('subscribing', () => (this.connected = false));
		sub.on('unsubscribed', () => (this.connected = false));
		sub.on('error', () => (this.connected = false));

		try {
			sub.subscribe();
			this.centrifuge.connect();
		} catch (error) {
			this.connected = false;
			toasterStore.showError(`real-time connection error\n${error}`);
		}
	}

	disconnect() {
		const { toasterStore } = stores;
		try {
			this.centrifuge?.disconnect();
		} catch (error) {
			this.connected = false;
			toasterStore.showError(`real-time connection error\n${error}`);
		}
	}
}
