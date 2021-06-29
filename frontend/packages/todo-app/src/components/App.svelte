<script>
  import { onMount, onDestroy } from "svelte";
  import { Toaster } from "todo-app-common";
  import { getTodos } from "../stores/todoStore";
  import { toasterStore } from "../stores/toasterStore";
  import { connect } from "../stores/notificationStore";
  import ConnectionStatus from "./ConnectionStatus.svelte";
  import Header from "./Header.svelte";
  import TodoForm from "./TodoForm.svelte";
  import Todos from "./Todos.svelte";
  import AppMenu from "./AppMenu.svelte";

  let disconnect = null;

  onMount(async () => {
    try {
      await getTodos();
    } catch (error) {
      toasterStore.showError(`todo loading failed\n${error}`);
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

<Toaster toasts={toasterStore.toasts} />
