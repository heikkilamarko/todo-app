<script>
  import { connected } from "../stores/notificationStore";
</script>

{#if $connected}
  <span class="text-success fw-bold">CONNECTED</span>
{:else}
  <span class="text-danger fw-bold">NO SIGNAL</span>
{/if}
