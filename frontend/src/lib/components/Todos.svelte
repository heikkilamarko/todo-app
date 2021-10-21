<script>
  import { isViewerRole } from "../shared/auth";
  import { stores } from "../stores";
  import Todo from "./Todo.svelte";
  import Empty from "./Empty.svelte";

  const isReadOnly = isViewerRole();

  const {
    todoStore: { todos, loading },
  } = stores;
</script>

{#each $todos as todo (todo.id)}
  <Todo {todo} {isReadOnly} />
{:else}
  {#if !$loading}
    <Empty />
  {/if}
{/each}
