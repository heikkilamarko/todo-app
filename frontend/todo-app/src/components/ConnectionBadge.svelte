<script>
  import { connected } from "../stores/notificationStore";
</script>

{#if $connected}
  <span class="badge rounded-pill bg-success">CONNECTED</span>
{:else}
  <span class="badge rounded-pill bg-danger">NO SIGNAL</span>
{/if}
