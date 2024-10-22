<script>
	import { createEventDispatcher, onDestroy } from 'svelte';
	import { fly } from 'svelte/transition';

	let { title, children } = $props();

	const dispatch = createEventDispatcher();
	const close = () => dispatch('close');

	let offcanvasEl = $state();

	function handleKeydown(e) {
		if (e.key === 'Escape') {
			close();
			return;
		}

		if (e.key === 'Tab') {
			const nodes = offcanvasEl.querySelectorAll('*');
			const tabbable = Array.from(nodes).filter((n) => n.tabIndex >= 0 && !n.disabled);

			let index = tabbable.indexOf(document.activeElement);
			if (index === -1 && e.shiftKey) index = 0;

			index += tabbable.length + (e.shiftKey ? -1 : 1);
			index %= tabbable.length;

			tabbable[index].focus();
			e.preventDefault();
		}
	}

	const previouslyFocused = document.activeElement;

	if (previouslyFocused) {
		onDestroy(() => {
			previouslyFocused.focus();
		});
	}
</script>

<svelte:window onkeydown={handleKeydown} />

<div
	class="offcanvas offcanvas-end show"
	role="dialog"
	aria-labelledby="offcanvas-label"
	bind:this={offcanvasEl}
	in:fly={{ x: 400, duration: 300 }}
	out:fly={{ x: 400, duration: 300 }}
>
	<div class="offcanvas-header">
		<h5 id="offcanvas-label" class="offcanvas-title">{title}</h5>
		<button type="button" class="btn-close text-reset" aria-label="Close" onclick={close}></button>
	</div>
	<div class="offcanvas-body">
		{@render children?.()}
	</div>
</div>

<style>
	.offcanvas {
		visibility: visible;
	}
</style>
