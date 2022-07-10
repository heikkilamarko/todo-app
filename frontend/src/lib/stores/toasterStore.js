import { writable } from 'svelte/store';

const TIMEOUT = 3000;

export default function createStore() {
	let toastId = 1;

	const toasts = writable([]);

	function showInfo(message) {
		showToast(message, 'info');
	}

	function showError(message) {
		showToast(message, 'error');
	}

	function showToast(message, type = 'info') {
		toasts.update((t) => [
			...t,
			{
				id: toastId++,
				message,
				type
			}
		]);

		setTimeout(() => toasts.update((t) => t.filter((_, i) => i)), TIMEOUT);
	}

	return { toasts, showInfo, showError, showToast };
}
