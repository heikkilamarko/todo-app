<script>
	import { page } from '$app/stores';
	import ErrorPage from './ErrorPage.svelte';
</script>

<ErrorPage status={$page.status} message={$page.error.message} />
