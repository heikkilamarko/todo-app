<script>
	let clazz;
	export { clazz as class };
	export let icon;
</script>

<span class={clazz}>
	<svelte:component this={icon} />
</span>
