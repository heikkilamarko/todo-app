<script>
	import { stores } from '$lib/shared/stores.js';
	import Todo from './Todo.svelte';
	import Empty from './Empty.svelte';

	const {
		user,
		todoStore: { todos, loading }
	} = stores;

	const isReadOnly = !user.hasPermission('todos.write');
</script>

{#each $todos as todo (todo.id)}
	<Todo {todo} {isReadOnly} />
{:else}
	{#if !$loading}
		<Empty />
	{/if}
{/each}
