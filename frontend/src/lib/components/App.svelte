<script>
	import { onMount, onDestroy } from 'svelte';
	import { isViewerRole } from '../shared/auth';
	import { stores, createStores } from '../stores';
	import Header from './Header.svelte';
	import ConnectionStatus from './ConnectionStatus.svelte';
	import TodoForm from './TodoForm.svelte';
	import Todos from './Todos.svelte';
	import AppMenu from './AppMenu.svelte';
	import Toaster from './Toaster.svelte';

	createStores();

	const isViewer = isViewerRole();

	const {
		toasterStore: { toasts, showError },
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

<Toaster {toasts} />
