<script>
	let { class: clazz, icon } = $props();

	const SvelteComponent = $derived(icon);
</script>

<span class={clazz}>
	<SvelteComponent />
</span>
