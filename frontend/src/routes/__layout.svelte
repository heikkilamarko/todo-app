<script context="module">
	import { startup } from '$lib/startup.js';

	export async function load() {
		try {
			await startup();
			return { status: 200 };
		} catch (error) {
			return { status: 500, error };
		}
	}
</script>

<script>
	import { onMount } from 'svelte';
	import { init } from '$lib/shared/auth.js';
	import { stores } from '$lib/shared/stores.js';
	import Toaster from '$lib/components/Toaster.svelte';

	const {
		toasterStore: { showError }
	} = stores;

	let isAuthenticated = false;

	onMount(async () => {
		try {
			isAuthenticated = await init();
		} catch (err) {
			showError(err);
		}
	});
</script>

{#if isAuthenticated}
	<slot />
{/if}

<Toaster />
