<script>
  import Todo from "./Todo.svelte";
  import { todos, complete } from "./stores/todoStore";
</script>

{#each $todos as todo (todo.id)}
  <div>
    <Todo {todo} on:click={() => complete(todo.id)} />
  </div>
{/each}
