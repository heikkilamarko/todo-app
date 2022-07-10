import { derived, writable } from 'svelte/store';

export default function createStore() {
	const name = writable('');
	const description = writable('');

	const closeOnCreate = writable(true);

	const todo = derived([name, description], ([$name, $description]) => ({
		name: $name,
		description: $description || null
	}));

	const isValid = derived(name, ($name) => !!$name);

	function reset() {
		name.set('');
		description.set('');
	}

	return { name, description, closeOnCreate, todo, isValid, reset };
}
