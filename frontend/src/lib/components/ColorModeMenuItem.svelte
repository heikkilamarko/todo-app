<script>
	import LightModeIcon from 'bootstrap-icons/icons/sun-fill.svg';
	import DarkModeIcon from 'bootstrap-icons/icons/moon-fill.svg';
	import SvgIcon from './SvgIcon.svelte';

	let colorMode = getColorMode();

	function getColorMode() {
		const storedColorMode = localStorage.getItem('app_color_mode');
		if (storedColorMode) return storedColorMode;
		return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
	}

	function setColorMode(mode) {
		localStorage.setItem('app_color_mode', mode);
		document.documentElement.setAttribute('data-bs-theme', mode);
	}

	function toggleColorMode() {
		colorMode = colorMode === 'light' ? 'dark' : 'light';
	}

	$: setColorMode(colorMode);
</script>

<li>
	<a class="dropdown-item" href="/" on:click|preventDefault={toggleColorMode}>
		{#if colorMode === 'light'}
			<SvgIcon icon={DarkModeIcon} class="text-primary pe-2" />
			Dark
		{:else}
			<SvgIcon icon={LightModeIcon} class="text-primary pe-2" />
			Light
		{/if}
	</a>
</li>
