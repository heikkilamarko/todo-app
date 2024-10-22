<script>
	import { fade } from 'svelte/transition';
	import { stores } from '$lib/shared/stores.js';

	const {
		notificationStore: { connected }
	} = stores;

	let show = $derived($connected === false);
</script>

{#if show}
	<section class="mb-4" transition:fade>
		<span class="badge bg-danger">NO SIGNAL</span>
	</section>
{/if}
