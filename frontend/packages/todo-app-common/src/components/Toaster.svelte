<script>
  import Toast from "./Toast.svelte";

  /** @type {import("../index").ToastsStore} */
  export let toasts;
</script>

<div class="a-toaster position-fixed start-0 end-0 bottom-0">
  {#each $toasts as toast (toast.id)}
    <Toast {toast} />
  {/each}
</div>

<style>
  .a-toaster {
    z-index: 2000;
  }
</style>
