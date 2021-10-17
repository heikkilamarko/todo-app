<script>
  import "../app.scss";
  import { onMount } from "svelte";
  import { load as loadConfig } from "$lib/shared/config";
  import { init as initAuth } from "$lib/shared/auth";
  import App from "$lib/components/App.svelte";
  import AppError from "$lib/components/AppError.svelte";

  let isReady = false;
  let error = null;

  onMount(async () => await import("bootstrap/js/dist/dropdown"));

  onMount(async () => {
    try {
      await loadConfig();
      const isAuthenticated = await initAuth();
      if (isAuthenticated) {
        isReady = true;
      }
    } catch (err) {
      console.log(err);
      error = err;
    }
  });
</script>

{#if isReady}
  <App />
{:else if error}
  <AppError {error} />
{/if}
