<script>
  import { fly } from "svelte/transition";

  export let notification;

  $: json = JSON.stringify(notification, null, 2);
</script>

<div
  class="p-2 my-4 shadow-sm rounded border"
  in:fly={{ x: 100, duration: 800, intro: true }}
>
  <pre>{json}</pre>
</div>

<style>
  pre {
    margin: 0;
  }
</style>
