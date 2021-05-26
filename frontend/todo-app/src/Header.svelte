<script>
  import { connected } from "./stores/notificationStore";
</script>

<h1 class="display-3 mt-2">
  Todo App
  {#if $connected}
    <span class="badge bg-success">CONNECTED</span>
  {:else}
    <span class="badge bg-danger">NO SIGNAL</span>
  {/if}
</h1>
