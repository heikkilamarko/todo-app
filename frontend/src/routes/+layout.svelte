<script>
	import '../app.scss';
	import 'bootstrap/js/dist/dropdown.js';
	import { onMount } from 'svelte';
	import { startup } from '$lib/startup.js';
	import Toaster from '$lib/components/Toaster.svelte';
	import ErrorPage from './ErrorPage.svelte';
	import AppMenu from '$lib/components/AppMenu.svelte';

	let { children } = $props();

	let status = $state(),
		message = $state();

	onMount(async () => {
		try {
			await startup();
			status = 200;
		} catch (err) {
			status = 500;
			message = err.message || 'An error occurred while starting the application';
		}
	});
</script>

{#if status === 200}
	<AppMenu />
	{@render children?.()}
	<Toaster />
{:else if status === 500}
	<ErrorPage {status} {message} />
{/if}
