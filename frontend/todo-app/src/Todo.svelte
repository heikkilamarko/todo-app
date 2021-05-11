<script>
  import { fly } from "svelte/transition";

  export let todo;
</script>

<div
  class="px-4 py-4 my-4 shadow-sm rounded border"
  in:fly={{ x: 100, duration: 600 }}
  out:fly={{ x: 500, duration: 300 }}
>
  <h1 class="display-6">{todo.name}</h1>
  <p class="text-muted">{todo.created_at.toLocaleString()}</p>
  <p class="pre">{todo.description || "-"}</p>
  <button type="button" class="btn btn-outline-success" on:click
    >Complete</button
  >
</div>

<style>
  .pre {
    white-space: pre-wrap;
  }
</style>
