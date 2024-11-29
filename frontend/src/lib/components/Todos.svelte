<script>
	import { stores } from '$lib/shared/stores.js';
	import Todo from './Todo.svelte';
	import Empty from './Empty.svelte';

	const { user, todoStore } = stores;

	const isReadOnly = !user.hasPermission('todo.write');
</script>

{#each todoStore.todos as todo (todo.id)}
	<Todo {todo} {isReadOnly} />
{:else}
	{#if !todoStore.loading}
		<Empty />
	{/if}
{/each}
