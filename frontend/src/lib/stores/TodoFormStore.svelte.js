export class TodoFormStore {
	name = $state('');
	description = $state('');
	closeOnCreate = $state(true);

	todo = $derived({
		name: this.name,
		description: this.description || null
	});

	isValid = $derived(!!this.name);

	reset() {
		this.name = '';
		this.description = '';
	}
}
