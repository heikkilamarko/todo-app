<script>
  import { onMount } from "svelte";
  import Notification from "./Notification.svelte";
  import TodoForm from "./TodoForm.svelte";
  import {
    getTodos,
    getSignalRConnection,
    toNotification,
    NOTIFICATION_METHOD_NAME,
  } from "./utils";

  let notifications = [];
  let isConnected = false;

  onMount(async () => {
    try {
      notifications = await getTodos();
    } catch (e) {
      alert(`Todo loading failed: ${e}`);
    }

    let connection = getSignalRConnection();

    connection.on(NOTIFICATION_METHOD_NAME, (notification) => {
      notifications = [toNotification(notification), ...notifications];
    });

    try {
      await connection.start();
      isConnected = true;
    } catch (e) {
      isConnected = false;
      console.error(e);
    }

    return () => connection.stop();
  });
</script>

<main class="container">
  <h1 class="display-3 mt-2">
    Todo App
    {#if isConnected}
      <span class="badge bg-success">CONNECTED</span>
    {/if}
  </h1>

  <section>
    <TodoForm />
  </section>

  <section>
    {#each notifications as notification (notification.id)}
      <div>
        <Notification {notification} />
      </div>
    {/each}
  </section>
</main>
