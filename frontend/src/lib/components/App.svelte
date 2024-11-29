<script>
	import { onMount, onDestroy } from 'svelte';
	import { stores } from '$lib/shared/stores.js';
	import Header from './Header.svelte';
	import ConnectionStatus from './ConnectionStatus.svelte';
	import TodoForm from './TodoForm.svelte';
	import Todos from './Todos.svelte';

	const { user, toasterStore, todoStore, notificationStore } = stores;

	const allowWrite = user.hasPermission('todo.write');

	onMount(async () => {
		try {
			await todoStore.getTodos();
		} catch (error) {
			toasterStore.showError(`todo loading failed\n${error}`);
		}

		await notificationStore.connect();
	});

	onDestroy(() => notificationStore.disconnect());
</script>

<main class="container">
	<Header>Todo App</Header>
	<ConnectionStatus />
	{#if allowWrite}
		<TodoForm />
	{/if}
	<Todos />
</main>
