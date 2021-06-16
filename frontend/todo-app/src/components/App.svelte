<script>
  import { onMount, onDestroy } from "svelte";
  import { showError } from "../stores/toasterStore";
  import { getTodos } from "../stores/todoStore";
  import { connect } from "../stores/notificationStore";
  import ConnectionStatus from "./ConnectionStatus.svelte";
  import Header from "./Header.svelte";
  import TodoForm from "./TodoForm.svelte";
  import Todos from "./Todos.svelte";
  import AppMenu from "./AppMenu.svelte";
  import Toaster from "./Toaster.svelte";

  let disconnect = null;

  onMount(async () => {
    try {
      await getTodos();
    } catch (error) {
      showError(`todo loading failed\n${error}`);
    }

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
