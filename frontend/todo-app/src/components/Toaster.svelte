<script>
  import { toasts } from "../stores/toasterStore";
  import Toast from "./Toast.svelte";
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
