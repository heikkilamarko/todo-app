<script>
	import { fly } from 'svelte/transition';
	import { stores } from '$lib/shared/stores.js';

	let { todo, isReadOnly = false } = $props();

	const { toasterStore, todoStore } = stores;

	async function handleComplete() {
		try {
			await todoStore.completeTodo(todo.id);
			toasterStore.showInfo('todo complete job started');
		} catch (error) {
			toasterStore.showError(`todo complete job failed\n${error}`);
		}
	}

	let createdAtLoc = $derived(todo.created_at.toLocaleString());
	let createdAtIso = $derived(todo.created_at.toISOString());

	let canComplete = $derived(!todoStore.loading);
</script>

<div class="card my-4" in:fly={{ x: 100, duration: 600 }} out:fly={{ x: 500, duration: 300 }}>
	<div class="card-body">
		<h4 class="card-title">{todo.name}</h4>
		<p class="text-muted" title={createdAtIso}>{createdAtLoc}</p>
		{#if todo.description}<p class="pre">{todo.description}</p>{/if}
		{#if !isReadOnly}
			<button
				type="button"
				class="btn btn-sm btn-outline-primary rounded-pill px-3"
				disabled={!canComplete}
				onclick={handleComplete}
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
