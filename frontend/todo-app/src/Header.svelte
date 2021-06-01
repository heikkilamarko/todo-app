<script>
  import { logout } from "./auth";
  import ConnectionBadge from "./ConnectionBadge.svelte";
</script>

<h1 class="display-3 mt-2">
  Todo App <ConnectionBadge />
  <button type="button" class="btn btn-danger float-end" on:click={logout}
    ><i class="bi bi-power" /> Logout</button
  >
</h1>
