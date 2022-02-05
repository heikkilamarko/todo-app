<script>
	import InfoCircleIcon from 'bootstrap-icons/icons/info-circle.svg';
	import GithubIcon from 'bootstrap-icons/icons/github.svg';
	import GridIcon from 'bootstrap-icons/icons/grid.svg';
	import PersonIcon from 'bootstrap-icons/icons/person.svg';
	import PersonFillIcon from 'bootstrap-icons/icons/person-fill.svg';
	import PowerIcon from 'bootstrap-icons/icons/power.svg';
	import SvgIcon from './SvgIcon.svelte';
	import { config } from '../shared/config';
	import { isViewerRole, logout, userName } from '../shared/auth';

	let isViewer = isViewerRole();

	let title = userName();
	if (isViewer) {
		title += ' (VIEWER)';
	}
</script>

<div class="btn-group position-fixed m-2 top-0 end-0">
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
		<li>
			<a class="dropdown-item" href="/about">
				<SvgIcon icon={InfoCircleIcon} class="text-primary pe-2" />
				About
			</a>
		</li>
		<li><hr class="dropdown-divider" /></li>
		<li>
			<a class="dropdown-item" href={config.profileUrl} target="_blank" rel="noreferrer">
				<SvgIcon icon={PersonIcon} class="text-primary pe-2" />
				Profile
			</a>
		</li>
		{#if !isViewer}
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
			<a class="dropdown-item" href="/" on:click|preventDefault={logout}>
				<SvgIcon icon={PowerIcon} class="text-primary pe-2" />
				Logout
			</a>
		</li>
	</ul>
</div>

<style>
	.dropdown-toggle::after {
		content: none;
	}
</style>
