<script>
  import Todo from "./Todo.svelte";
  import { todos, completeTodo } from "../stores/todoStore";
</script>

{#each $todos as todo (todo.id)}
  <Todo {todo} on:click={() => completeTodo(todo.id)} />
{/each}
