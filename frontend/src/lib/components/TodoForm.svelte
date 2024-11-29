<script>
	import { preventDefault } from 'svelte/legacy';
	import Offcanvas from './Offcanvas.svelte';
	import { stores } from '$lib/shared/stores.js';

	const { toasterStore, todoFormStore, todoStore } = stores;

	let showOffcanvas = $state(false);

	async function handleSubmit() {
		try {
			await todoStore.createTodo(todoFormStore.todo);
			toasterStore.showInfo('todo create job started');
			todoFormStore.reset();
			if (todoFormStore.closeOnCreate) {
				closeOffcanvas();
			}
		} catch (error) {
			toasterStore.showError(`todo create job failed\n${error}`);
		}
	}

	function closeOffcanvas() {
		showOffcanvas = false;
	}

	function toggleOffcanvas() {
		showOffcanvas = !showOffcanvas;
	}

	let canCreate = $derived(todoFormStore.isValid && !todoStore.loading);
</script>

<button class="btn btn-primary rounded-pill px-3" type="button" onclick={toggleOffcanvas}>
	New Todo
</button>

{#if showOffcanvas}
	<Offcanvas title="New Todo" on:close={closeOffcanvas}>
		<form spellcheck="false" autocomplete="off" onsubmit={preventDefault(handleSubmit)}>
			<div class="mb-3">
				<label for="name" class="form-label">Name <span class="text-danger">*</span></label>
				<input
					type="text"
					class="form-control"
					id="name"
					placeholder="Name..."
					bind:value={todoFormStore.name}
				/>
			</div>
			<div class="mb-3">
				<label for="description" class="form-label">Description</label>
				<textarea
					class="form-control"
					id="description"
					rows="5"
					bind:value={todoFormStore.description}
				></textarea>
			</div>
			<div class="mb-3">
				<div class="form-check">
					<input
						id="close-on-save-check"
						class="form-check-input"
						type="checkbox"
						bind:checked={todoFormStore.closeOnCreate}
					/>
					<label class="form-check-label" for="close-on-save-check"> Close on Create </label>
				</div>
			</div>
			<button type="submit" class="btn btn-primary rounded-pill px-3" disabled={!canCreate}>
				Create
			</button>
		</form>
	</Offcanvas>
{/if}
