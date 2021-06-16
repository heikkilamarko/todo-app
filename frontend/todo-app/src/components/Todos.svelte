<script>
  import { showInfo, showError } from "../stores/toasterStore";
  import { todos, completeTodo } from "../stores/todoStore";
  import Todo from "./Todo.svelte";
  import Empty from "./Empty.svelte";

  async function complete(id) {
    try {
      await completeTodo(id);
      showInfo("todo complete job started");
    } catch (error) {
      showError(`todo complete job failed\n${error}`);
    }
  }
</script>

{#each $todos as todo (todo.id)}
  <Todo {todo} on:click={() => complete(todo.id)} />
{:else}
  <Empty />
{/each}
