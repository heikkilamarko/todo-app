<script>
  import { onMount } from "svelte";
  import { fade, fly } from "svelte/transition";
  import { backOut } from "svelte/easing";

  let toasts = [];
  let timeout = 2500;
  let nextId = 0;

  function pushToast(message, type = "success") {
    toasts = [
      ...toasts,
      {
        id: ++nextId,
        message,
        type,
      },
    ];

    setTimeout(() => (toasts = toasts.filter((a, i) => i)), timeout);
  }

  onMount(() => {
    window.pushToast = pushToast;
  });
</script>

<div class="a-toaster position-fixed start-0 end-0 bottom-0">
  {#each toasts as toast (toast.id)}
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
  {/each}
</div>

<style>
  .a-toaster {
    z-index: 2000;
  }
  .a-toast {
    white-space: pre-wrap;
  }
</style>
