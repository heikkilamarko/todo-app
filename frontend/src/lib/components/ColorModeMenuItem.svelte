<script>
	let colorMode = $state(getColorMode());

	$effect(() => {
		setColorMode(colorMode);
	});

	function getColorMode() {
		const storedColorMode = localStorage.getItem('app_color_mode');
		if (storedColorMode) return storedColorMode;
		return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
	}

	function setColorMode(mode) {
		localStorage.setItem('app_color_mode', mode);
		document.documentElement.setAttribute('data-bs-theme', mode);
	}

	function toggleColorMode(e) {
		e.preventDefault();
		colorMode = colorMode === 'light' ? 'dark' : 'light';
	}
</script>

<li>
	<a class="dropdown-item" href="/" onclick={toggleColorMode}>
		{#if colorMode === 'light'}
			<span class="icon--bi icon--bi--moon-fill text-primary pe-2"></span>
			Dark
		{:else}
			<span class="icon--bi icon--bi--sun-fill text-primary pe-2"></span>
			Light
		{/if}
	</a>
</li>
