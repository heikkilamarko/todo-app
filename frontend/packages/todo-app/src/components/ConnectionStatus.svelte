<script>
  import { stores } from "../stores";

  const {
    notificationStore: { connected },
  } = stores;

  $: status = $connected ? "CONNECTED" : "NO SIGNAL";
</script>

<div class="text-primary fw-bold my-1">
  <i class="bi" class:bi-check-lg={$connected} class:bi-x-lg={!$connected} />
  {status}
</div>
