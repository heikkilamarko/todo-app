<script>
  import { fly } from "svelte/transition";

  export let todo;
</script>

<div
  class="px-4 py-2 my-4 shadow-sm rounded border"
  in:fly={{ x: 100, duration: 800 }}
>
  <dl>
    <dt>Name</dt>
    <dd>{todo.name}</dd>
    <dt>Description</dt>
    <dd class="pre">{todo.description || "-"}</dd>
    <dt>Created</dt>
    <dd>{todo.created_at.toLocaleString()}</dd>
  </dl>
</div>

<style>
  dl {
    margin: 0;
  }

  dd:last-child {
    margin-bottom: 0;
  }

  .pre {
    white-space: pre-wrap;
  }
</style>
