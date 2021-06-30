<script>
  import { fly } from "svelte/transition";
  import { CheckLgIcon } from "todo-app-common";
  import { stores } from "../stores";

  /** @type {import("../types").Todo} */
  export let todo;

  const {
    toasterStore: { showInfo, showError },
    todoStore: { loading, completeTodo },
  } = stores;

  async function handleComplete() {
    try {
      await completeTodo(todo.id);
      showInfo("todo complete job started");
    } catch (error) {
      showError(`todo complete job failed\n${error}`);
    }
  }

  $: createdAtLoc = todo.created_at.toLocaleString();
  $: createdAtIso = todo.created_at.toISOString();

  $: canComplete = !$loading;
</script>

<div
  class="px-4 py-4 my-4 bg-white shadow-sm rounded border"
  in:fly={{ x: 100, duration: 600 }}
  out:fly={{ x: 500, duration: 300 }}
>
  <h1 class="display-6 text-primary">{todo.name}</h1>
  <p class="text-muted" title={createdAtIso}>{createdAtLoc}</p>
  <p class="pre">{todo.description || "-"}</p>
  <button
    type="button"
    class="btn btn-outline-primary rounded-pill px-3"
    disabled={!canComplete}
    on:click={handleComplete}
  >
    <CheckLgIcon />
    Complete
  </button>
</div>

<style>
  .pre {
    white-space: pre-wrap;
  }
</style>
