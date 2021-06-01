<script>
  import Todo from "./Todo.svelte";
  import { todos, complete } from "./stores/todoStore";
</script>

{#each $todos as todo (todo.id)}
  <Todo {todo} on:click={() => complete(todo.id)} />
{/each}
