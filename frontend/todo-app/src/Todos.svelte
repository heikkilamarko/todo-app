<script>
  import Todo from "./Todo.svelte";
  import { completeTodo, showError, showInfo } from "./utils";

  export let todos = [];

  async function complete(id) {
    try {
      await completeTodo(id);
      showInfo("Todo successfully sent for processing");
    } catch (e) {
      showError(`Todo completion failed\n${e}`);
    }
  }
</script>

{#each todos as todo (todo.id)}
  <div>
    <Todo {todo} on:click={() => complete(todo.id)} />
  </div>
{/each}
