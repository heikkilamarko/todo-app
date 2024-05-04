import { initializeFaro } from '@grafana/faro-web-sdk';
import { stores } from './stores.js';

export default function logging() {
	return initializeFaro({
		url: stores.config.loggingUrl,
		app: {
			name: 'todo-app',
			version: '1.0.0'
		}
	});
}
