<script>
  import { onMount } from "svelte";
  import Todo from "./Todo.svelte";
  import TodoForm from "./TodoForm.svelte";
  import {
    getTodos,
    getSignalRConnection,
    toTodo,
    NOTIFICATION_METHOD_NAME,
  } from "./utils";

  let todos = [];
  let isConnected = false;

  onMount(async () => {
    try {
      todos = await getTodos();
    } catch (e) {
      alert(`Todo loading failed: ${e}`);
    }

    let connection = getSignalRConnection();

    connection.onclose(() => (isConnected = false));
    connection.onreconnecting(() => (isConnected = false));
    connection.onreconnected(() => (isConnected = true));

    connection.on(NOTIFICATION_METHOD_NAME, (notification) => {
      const { type, data } = notification ?? {};

      if (type === "todo.created.ok") {
        todos = [toTodo(data), ...todos];
        window?.confetti();
      } else if (type === "todo.completed.ok") {
        todos = todos.filter((t) => t.id !== data.id);
      } else {
        console.log("unknown notification received", type, data);
      }
    });

    try {
      await connection.start();
      isConnected = true;
    } catch (e) {
      alert(`Real-time connection failed: ${e}`);
    }

    return () => connection.stop();
  });
</script>

<main class="container">
  <h1 class="display-3 mt-2">
    Todo App
    {#if isConnected}
      <span class="badge bg-success">CONNECTED</span>
    {:else}
      <span class="badge bg-danger">NO SIGNAL</span>
    {/if}
  </h1>

  <section>
    <TodoForm />
  </section>

  <section>
    {#each todos as todo (todo.id)}
      <div>
        <Todo {todo} />
      </div>
    {/each}
  </section>
</main>
