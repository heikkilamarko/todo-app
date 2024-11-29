import { stores } from './shared/stores.js';
import config from './shared/config.js';
import logging from './shared/logging.js';
import auth from './shared/auth.js';
import user from './shared/user.js';
import { ToasterStore } from './stores/ToasterStore.svelte.js';
import { TodoStore } from './stores/TodoStore.svelte.js';
import { TodoFormStore } from './stores/TodoFormStore.svelte.js';
import { NotificationStore } from './stores/NotificationStore.svelte.js';

let startupOk = false;

export async function startup() {
	if (startupOk) return;
	await initStores();
	startupOk = true;
	console.info('app startup ready');
}

async function initStores() {
	try {
		stores.config = await config();
		stores.logging = logging();
		stores.auth = await auth();
		stores.user = await user();
		stores.toasterStore = new ToasterStore();
		stores.todoStore = new TodoStore();
		stores.todoFormStore = new TodoFormStore();
		stores.notificationStore = new NotificationStore();
	} catch (err) {
		console.error(err);
		throw new Error('app startup error');
	}
}
