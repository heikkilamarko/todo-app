<script>
  import { onMount } from "svelte";
  import Header from "./Header.svelte";
  import TodoForm from "./TodoForm.svelte";
  import Todos from "./Todos.svelte";
  import Toaster from "./Toaster.svelte";
  import { load as loadConfig } from "./stores/configStore";
  import { load as loadTodos } from "./stores/todoStore";
  import { connect } from "./stores/notificationStore";

  onMount(async () => {
    await loadConfig();
    await loadTodos();
    return await connect();
  });
</script>

<main class="container">
  <Header />

  <section>
    <TodoForm />
  </section>

  <section>
    <Todos />
  </section>
</main>

<Toaster />
