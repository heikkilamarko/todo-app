<script>
  import "../app.scss";
  import "bootstrap/js/dist/dropdown";
  import { onMount } from "svelte";
  import { load as loadConfig } from "$lib/shared/config";
  import { init as initAuth } from "$lib/shared/auth";
  import AppError from "$lib/components/AppError.svelte";

  let isAuthenticated = false;
  let error = null;

  onMount(async () => {
    try {
      await loadConfig();
      isAuthenticated = await initAuth();
    } catch (err) {
      console.log(err);
      error = err;
    }
  });
</script>

{#if error}
  <AppError {error} />
{:else if isAuthenticated}
  <slot />
{/if}
