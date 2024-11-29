const TIMEOUT = 3000;

let toastId = 1;

export class ToasterStore {
	toasts = $state([]);

	showInfo(message) {
		this.showToast(message, 'info');
	}

	showError(message) {
		this.showToast(message, 'error');
	}

	showToast(message, type = 'info') {
		this.toasts.push({
			id: toastId++,
			message,
			type
		});

		setTimeout(() => (this.toasts = this.toasts.filter((_, i) => i)), TIMEOUT);
	}
}
