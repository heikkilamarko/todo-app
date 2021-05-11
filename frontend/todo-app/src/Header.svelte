<script>
  export let isConnected = false;
</script>

<h1 class="display-3 mt-2">
  Todo App
  {#if isConnected}
    <span class="badge bg-success">CONNECTED</span>
  {:else}
    <span class="badge bg-danger">NO SIGNAL</span>
  {/if}
</h1>
