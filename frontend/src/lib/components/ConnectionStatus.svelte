<script>
	import { fade } from 'svelte/transition';
	import { stores } from '$lib/shared/stores.js';

	const { notificationStore } = stores;
</script>

{#if !notificationStore.connected}
	<section class="mb-4" transition:fade>
		<span class="badge bg-danger">NO SIGNAL</span>
	</section>
{/if}
