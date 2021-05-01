<script>
  import { onMount } from "svelte";
  import {
    getTodos,
    getSignalRConnection,
    SIGNALR_RECEIVE_NOTIFICATION,
  } from "./utils";

  let notifications = [];
  let isConnected = false;

  onMount(async () => {
    notifications = await getTodos();

    let connection = getSignalRConnection();

    connection.on(SIGNALR_RECEIVE_NOTIFICATION, (data) => {
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

  $: connection = isConnected ? "CONNECTED" : "NO CONNECTION";
  $: formattedNotifications = notifications.map((n) =>
    JSON.stringify(n, null, 2)
  );
</script>

<main class="container">
  <h1 class="display-3 mt-2">
    Todo App <span class="badge bg-secondary" class:bg-success={isConnected}
      >{connection}</span
    >
  </h1>

  <section>
    {#each formattedNotifications as notification}
      <div class="p-2 my-4 shadow-sm rounded border">
        <pre>
        {notification}
      </pre>
      </div>
    {/each}
  </section>
</main>

<style>
  pre {
    margin: 0;
  }
</style>
