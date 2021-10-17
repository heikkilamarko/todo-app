<script>
  import { onMount, onDestroy } from "svelte";
  import { stores, createStores } from "../stores";
  import Toaster from "./Toaster.svelte";
  import Header from "./Header.svelte";
  import ConnectionStatus from "./ConnectionStatus.svelte";
  import TodoForm from "./TodoForm.svelte";
  import Todos from "./Todos.svelte";
  import AppMenu from "./AppMenu.svelte";

  createStores();

  const {
    toasterStore: { toasts, showError },
    todoStore: { getTodos },
    notificationStore: { connect },
  } = stores;

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
  <Header />
  <ConnectionStatus />
  <TodoForm />
  <Todos />
</main>

<AppMenu />

<Toaster {toasts} />
