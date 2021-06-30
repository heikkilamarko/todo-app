<script>
  import { CheckLgIcon, XLgIcon } from "todo-app-common";
  import { stores } from "../stores";

  const {
    notificationStore: { connected },
  } = stores;

  $: status = $connected ? "CONNECTED" : "NO SIGNAL";
</script>

<div class="text-primary fw-bold my-1">
  {#if $connected} <CheckLgIcon /> {:else} <XLgIcon /> {/if}
  {status}
</div>
