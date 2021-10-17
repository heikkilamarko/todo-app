<script>
  import { stores } from "../stores";
  import Todo from "./Todo.svelte";
  import Empty from "./Empty.svelte";

  const {
    todoStore: { todos, loading },
  } = stores;
</script>

{#each $todos as todo (todo.id)}
  <Todo {todo} />
{:else}
  {#if !$loading}
    <Empty />
  {/if}
{/each}
