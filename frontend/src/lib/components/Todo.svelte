<script>
	import { fly } from 'svelte/transition';
	import { stores } from '$lib/shared/stores.js';

	export let todo;

	export let isReadOnly = false;

	const {
		toasterStore: { showInfo, showError },
		todoStore: { loading, completeTodo }
	} = stores;

	async function handleComplete() {
		try {
			await completeTodo(todo.id);
			showInfo('todo complete job started');
		} catch (error) {
			showError(`todo complete job failed\n${error}`);
		}
	}

	$: createdAtLoc = todo.created_at.toLocaleString();
	$: createdAtIso = todo.created_at.toISOString();

	$: canComplete = !$loading;
</script>

<div
	class="card my-4"
	in:fly|local={{ x: 100, duration: 600 }}
	out:fly|local={{ x: 500, duration: 300 }}
>
	<div class="card-body">
		<h4 class="card-title">{todo.name}</h4>
		<p class="text-muted" title={createdAtIso}>{createdAtLoc}</p>
		{#if todo.description}<p class="pre">{todo.description}</p>{/if}
		{#if !isReadOnly}
			<button
				type="button"
				class="btn btn-sm btn-outline-primary rounded-pill px-3"
				disabled={!canComplete}
				on:click={handleComplete}
			>
				Complete
			</button>
		{/if}
	</div>
</div>

<style>
	.pre {
		white-space: pre-wrap;
	}
</style>
