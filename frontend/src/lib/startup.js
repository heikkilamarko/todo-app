import { stores } from './shared/stores.js';
import config from './shared/config.js';
import logging from './shared/logging.js';
import auth from './shared/auth.js';
import user from './shared/user.js';
import toasterStore from './stores/toasterStore.js';
import todoStore from './stores/todoStore.js';
import todoFormStore from './stores/todoFormStore.js';
import notificationStore from './stores/notificationStore.js';

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
		stores.toasterStore = toasterStore();
		stores.todoStore = todoStore();
		stores.todoFormStore = todoFormStore();
		stores.notificationStore = notificationStore();
	} catch (err) {
		console.error(err);
		throw new Error('app startup error');
	}
}
