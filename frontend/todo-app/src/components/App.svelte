<script>
  import { onMount, onDestroy } from "svelte";
  import Header from "./Header.svelte";
  import TodoForm from "./TodoForm.svelte";
  import Todos from "./Todos.svelte";
  import LogoutButton from "./LogoutButton.svelte";
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

<main class="container position-relative">
  <Header />
  <TodoForm />
  <Todos />
</main>

<LogoutButton />

<Toaster />
