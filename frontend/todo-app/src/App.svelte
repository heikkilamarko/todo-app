<script>
  import { onMount } from "svelte";
  import Toaster from "./Toaster.svelte";
  import Header from "./Header.svelte";
  import TodoForm from "./TodoForm.svelte";
  import Todos from "./Todos.svelte";
  import { Notification } from "./constants";
  import {
    getTodos,
    getSignalRConnection,
    toTodo,
    showError,
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
        case Notification.TodoCompletedOk:
          await load();
          break;
        case Notification.TodoCreatedError:
        case Notification.TodoCompletedError:
          showError(`ERROR: ${data.code}\n${data.message || "-"}`);
          break;
      }
    });

    try {
      await connection.start();
      isConnected = true;
    } catch (e) {
      showError(`Real-time connection failed\n${e}`);
    }

    async function load() {
      try {
        todos = await getTodos();
      } catch (e) {
        showError(`Todo loading failed\n${e}`);
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

<Toaster />
