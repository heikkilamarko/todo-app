<script>
  import { onMount } from "svelte";
  import Notification from "./Notification.svelte";
  import {
    getTodos,
    getSignalRConnection,
    NOTIFICATION_METHOD_NAME,
  } from "./utils";

  let notifications = [];
  let isConnected = false;

  onMount(async () => {
    notifications = await getTodos();

    let connection = getSignalRConnection();

    connection.on(NOTIFICATION_METHOD_NAME, (data) => {
      notifications = [data, ...notifications];
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
    {#each notifications as notification (notification.id)}
      <div>
        <Notification {notification} />
      </div>
    {/each}
  </section>
</main>
