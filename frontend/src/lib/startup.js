import 'bootstrap/js/dist/dropdown';
import { stores } from './shared/stores.js';
import { config } from './shared/config.js';
import { initAuth } from './shared/auth.js';
import toasterStore from './stores/toasterStore.js';
import todoStore from './stores/todoStore.js';
import todoFormStore from './stores/todoFormStore.js';
import notificationStore from './stores/notificationStore.js';

let startupOk = false;

export async function startup() {
	if (startupOk) return;
	await loadConfig();
	await initAuth();
	initStores();
	startupOk = true;
}

async function loadConfig() {
	try {
		stores.config = await config();
	} catch (err) {
		console.error(err);
		throw new Error('An error occurred while loading application configuration');
	}
}

function initStores() {
	try {
		stores.toasterStore = toasterStore();
		stores.todoStore = todoStore();
		stores.todoFormStore = todoFormStore();
		stores.notificationStore = notificationStore();
	} catch (err) {
		console.error(err);
		throw new Error('An error occurred while starting the application');
	}
}
