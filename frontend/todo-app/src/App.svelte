<script>
  import { onMount } from "svelte";
  import Header from "./Header.svelte";
  import TodoForm from "./TodoForm.svelte";
  import Todos from "./Todos.svelte";
  import { Notification } from "./constants";
  import {
    getTodos,
    getSignalRConnection,
    toTodo,
    NOTIFICATION_METHOD_NAME,
  } from "./utils";

  let todos = [];
  let isConnected = false;

  onMount(async () => {
    await load();

    let connection = getSignalRConnection();

    connection.onclose(() => (isConnected = false));
    connection.onreconnecting(() => (isConnected = false));
    connection.onreconnected(() => (isConnected = true));

    connection.on(NOTIFICATION_METHOD_NAME, async (notification) => {
      const { type, data } = notification ?? {};
      switch (type) {
        case Notification.TodoCreatedOk:
          await load();
          window?.confetti();
          break;
        case Notification.TodoCompletedOk:
          await load();
          break;
        case Notification.TodoCreatedError:
        case Notification.TodoCompletedError:
          // TODO: Implement proper error handling
          alert(JSON.stringify(data, null, 2));
          break;
        default:
          console.log("unknown notification received", type, data);
      }
    });

    try {
      await connection.start();
      isConnected = true;
    } catch (e) {
      alert(`Real-time connection failed: ${e}`);
    }

    async function load() {
      try {
        todos = await getTodos();
      } catch (e) {
        alert(`Todo loading failed: ${e}`);
      }
    }

    return () => connection.stop();
  });
</script>

<main class="container">
  <Header {isConnected} />

  <section>
    <TodoForm />
  </section>

  <section>
    <Todos {todos} />
  </section>
</main>
