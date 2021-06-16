<script>
  import { fade, fly } from "svelte/transition";
  import { backOut } from "svelte/easing";

  export let toast;
</script>

<div
  class="a-toast bg-{toast.type} text-white rounded w-75 mx-auto my-2 px-4 py-2"
  in:fly={{
    delay: 0,
    duration: 300,
    x: 0,
    y: 50,
    opacity: 0.1,
    easing: backOut,
  }}
  out:fade={{
    duration: 500,
    opacity: 0,
  }}
>
  {toast.message}
</div>

<style>
  .a-toast {
    white-space: pre-wrap;
  }
</style>
