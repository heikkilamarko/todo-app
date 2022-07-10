<script>
	import { fade, fly } from 'svelte/transition';
	import { backOut } from 'svelte/easing';

	export let toast;

	function getToastType(type) {
		switch (type) {
			case 'info':
				return 'primary';
			case 'error':
				return 'danger';
			default:
				return 'primary';
		}
	}

	$: toastType = getToastType(toast.type);
</script>

<div
	class="a-toast bg-{toastType} text-white rounded w-75 mx-auto my-2 px-4 py-2"
	in:fly|local={{
		delay: 0,
		duration: 300,
		x: 0,
		y: 50,
		opacity: 0.1,
		easing: backOut
	}}
	out:fade|local={{
		duration: 500
	}}
>
	{toast.message}
</div>

<style>
	.a-toast {
		white-space: pre-wrap;
	}
</style>
