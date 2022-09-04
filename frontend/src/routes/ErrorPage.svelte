<script>
	export let status = 500;
	export let message = 'An error occurred while running the application';
</script>

<main class="container py-3">
	<section
		class="bg-danger text-white rounded shadow p-4 d-flex flex-column align-items-center justify-content-center"
	>
		<h1 class="display-2 fw-lighter">Oops!</h1>
		<p class="display-1">
			{status}
		</p>
		<p class="fs-4 fw-light">
			{message}
		</p>
	</section>
</main>
