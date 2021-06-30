<script>
  import { CheckLg, XLg } from "todo-app-common";
  import { stores } from "../stores";

  const {
    notificationStore: { connected },
  } = stores;

  $: status = $connected ? "CONNECTED" : "NO SIGNAL";
</script>

<div class="text-primary fw-bold my-1">
  {#if $connected} <CheckLg /> {:else} <XLg /> {/if}
  {status}
</div>
