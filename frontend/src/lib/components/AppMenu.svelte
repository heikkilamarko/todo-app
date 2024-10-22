<script>
	import { preventDefault } from 'svelte/legacy';
	import InfoCircleIcon from 'bootstrap-icons/icons/info-circle.svg';
	import GithubIcon from 'bootstrap-icons/icons/github.svg';
	import GridIcon from 'bootstrap-icons/icons/grid.svg';
	import PersonIcon from 'bootstrap-icons/icons/person.svg';
	import PersonFillIcon from 'bootstrap-icons/icons/person-fill.svg';
	import PowerIcon from 'bootstrap-icons/icons/power.svg';
	import { stores } from '$lib/shared/stores.js';
	import SvgIcon from './SvgIcon.svelte';
	import ColorModeMenuItem from './ColorModeMenuItem.svelte';

	const { auth, config, user } = stores;

	let allowWrite = user.hasPermission('todo.write');

	let title = user.username();
</script>

<div class="btn-group position-fixed z-3 m-2 top-0 end-0">
	<button
		type="button"
		class="btn btn-outline-primary dropdown-toggle rounded-pill px-3"
		data-bs-toggle="dropdown"
		aria-expanded="false"
	>
		<PersonFillIcon />
		{title}
	</button>
	<ul class="dropdown-menu">
		<ColorModeMenuItem />
		<li><hr class="dropdown-divider" /></li>
		<li>
			<a class="dropdown-item" href={config.profileUrl} target="_blank" rel="noreferrer">
				<SvgIcon icon={PersonIcon} class="text-primary pe-2" />
				Profile
			</a>
		</li>
		{#if allowWrite}
			<li>
				<a class="dropdown-item" href={config.dashboardUrl} target="_blank" rel="noreferrer">
					<SvgIcon icon={GridIcon} class="text-primary pe-2" />
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
				<SvgIcon icon={GithubIcon} class="text-primary pe-2" />
				GitHub
			</a>
		</li>
		<li><hr class="dropdown-divider" /></li>
		<li>
			<a class="dropdown-item" href="/about">
				<SvgIcon icon={InfoCircleIcon} class="text-primary pe-2" />
				About
			</a>
		</li>
		<li>
			<a class="dropdown-item" href="/" onclick={preventDefault(() => auth.signOut())}>
				<SvgIcon icon={PowerIcon} class="text-primary pe-2" />
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
