<script>
	import { stores } from '$lib/shared/stores.js';
	import ColorModeMenuItem from './ColorModeMenuItem.svelte';

	const { auth, config, user } = stores;

	let allowWrite = user.hasPermission('todo.write');

	let title = user.username();

	function handleSignOut(e) {
		e.preventDefault();
		auth.signOut();
	}
</script>

<div class="btn-group position-fixed z-3 m-2 top-0 end-0">
	<button
		type="button"
		class="btn btn-outline-primary dropdown-toggle rounded-pill px-3"
		data-bs-toggle="dropdown"
		aria-expanded="false"
	>
		<span class="bi--icon bi--person-fill pe-2"></span>
		{title}
	</button>
	<ul class="dropdown-menu">
		<ColorModeMenuItem />
		<li><hr class="dropdown-divider" /></li>
		<li>
			<a class="dropdown-item" href={config.profileUrl} target="_blank" rel="noreferrer">
				<span class="bi--icon bi--person text-primary pe-2"></span>
				Profile
			</a>
		</li>
		{#if allowWrite}
			<li>
				<a class="dropdown-item" href={config.dashboardUrl} target="_blank" rel="noreferrer">
					<span class="bi--icon bi--grid text-primary pe-2"></span>
					Dashboard
				</a>
			</li>
		{/if}
		<li>
			<a
				class="dropdown-item"
				href="https://github.com/heikkilamarko/todo-app"
				target="_blank"
				rel="noreferrer"
			>
				<span class="bi--icon bi--github text-primary pe-2"></span>
				GitHub
			</a>
		</li>
		<li><hr class="dropdown-divider" /></li>
		<li>
			<a class="dropdown-item" href="/about">
				<span class="bi--icon bi--info-circle text-primary pe-2"></span>
				About
			</a>
		</li>
		<li>
			<a class="dropdown-item" href="/" onclick={handleSignOut}>
				<span class="bi--icon bi--power text-primary pe-2"></span>
				Sign out
			</a>
		</li>
	</ul>
</div>

<style>
	.dropdown-toggle::after {
		content: none;
	}
</style>
