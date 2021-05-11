<script>
  import Todo from "./Todo.svelte";

  export let todos = [];
</script>

{#each todos as todo (todo.id)}
  <div>
    <Todo {todo} />
  </div>
{/each}
