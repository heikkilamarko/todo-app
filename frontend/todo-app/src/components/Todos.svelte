<script>
  import { todos, loading } from "../stores/todoStore";
  import Todo from "./Todo.svelte";
  import Empty from "./Empty.svelte";
</script>

{#each $todos as todo (todo.id)}
  <Todo {todo} />
{:else}
  {#if !$loading}
    <Empty />
  {/if}
{/each}
