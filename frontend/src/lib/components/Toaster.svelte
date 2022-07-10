<script>
	import { stores } from '$lib/shared/stores.js';
	import Toast from './Toast.svelte';

	const {
		toasterStore: { toasts }
	} = stores;
</script>

<div class="position-fixed start-0 end-0 bottom-0 app-toaster">
	<div class="container">
		{#each $toasts as toast (toast.id)}
			<Toast {toast} />
		{/each}
	</div>
</div>

<style>
	.app-toaster {
		z-index: 2000;
	}
</style>
