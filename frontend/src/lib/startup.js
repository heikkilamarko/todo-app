import 'bootstrap/js/dist/dropdown';
import { stores } from './shared/stores.js';
import { config } from './shared/config.js';
import toasterStore from './stores/toasterStore.js';
import todoStore from './stores/todoStore.js';
import todoFormStore from './stores/todoFormStore.js';
import notificationStore from './stores/notificationStore.js';
import '../app.scss';

export async function startup() {
	await loadConfig();
	initStores();
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
