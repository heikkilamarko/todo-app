<script>
	import { onMount, onDestroy } from 'svelte';
	import { isViewerRole } from '$lib/shared/auth.js';
	import { stores } from '$lib/shared/stores.js';
	import AppMenu from './AppMenu.svelte';
	import Header from './Header.svelte';
	import ConnectionStatus from './ConnectionStatus.svelte';
	import TodoForm from './TodoForm.svelte';
	import Todos from './Todos.svelte';

	const isViewer = isViewerRole();

	const {
		toasterStore: { showError },
		todoStore: { getTodos },
		notificationStore: { connect }
	} = stores;

	let disconnect = null;

	onMount(async () => {
		try {
			await getTodos();
		} catch (error) {
			showError(`todo loading failed\n${error}`);
		}

		disconnect = await connect();
	});

	onDestroy(() => disconnect?.());
</script>

<AppMenu />

<main class="container">
	<Header>Todo App</Header>
	<ConnectionStatus />
	{#if !isViewer}
		<TodoForm />
	{/if}
	<Todos />
</main>
