<script>
	import { fade } from 'svelte/transition';
	import { stores } from '../stores';

	const {
		notificationStore: { connected }
	} = stores;

	$: show = $connected === false;
</script>

{#if show}
	<section class="mb-4" transition:fade>
		<span class="badge bg-danger">NO SIGNAL</span>
	</section>
{/if}
