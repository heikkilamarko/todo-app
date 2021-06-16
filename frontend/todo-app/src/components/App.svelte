<script>
  import { onMount, onDestroy } from "svelte";
  import ConnectionStatus from "./ConnectionStatus.svelte";
  import Header from "./Header.svelte";
  import TodoForm from "./TodoForm.svelte";
  import Todos from "./Todos.svelte";
  import AppMenu from "./AppMenu.svelte";
  import Toaster from "./Toaster.svelte";
  import { load } from "../stores/todoStore";
  import { connect } from "../stores/notificationStore";

  let disconnect = null;

  onMount(async () => {
    await load();
    disconnect = await connect();
  });

  onDestroy(() => disconnect?.());
</script>

<main class="container">
  <ConnectionStatus />
  <Header />
  <TodoForm />
  <Todos />
</main>

<AppMenu />

<Toaster />
