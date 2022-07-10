<script context="module">
	export function load({ error, status }) {
		return { props: { status, error } };
	}
</script>

<script>
	export let status;
	export let error;

	$: message = error?.message || 'An error occurred';
</script>

<main class="container py-3">
	<section
		class="bg-danger text-white rounded shadow p-4 d-flex flex-column align-items-center justify-content-center"
	>
		<h1 class="display-2 fw-lighter">Oops!</h1>
		<p class="display-1">
			{status || 500}
		</p>
		<p class="fs-4 fw-light">
			{message}
		</p>
	</section>
</main>
