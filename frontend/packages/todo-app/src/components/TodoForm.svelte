<script>
  import { Offcanvas } from "todo-app-common";
  import { stores } from "../stores";

  const {
    toasterStore: { showInfo, showError },
    todoFormStore: { name, description, closeOnCreate, todo, isValid, reset },
    todoStore: { loading, createTodo },
  } = stores;

  let showOffcanvas = false;

  async function handleSubmit() {
    try {
      await createTodo($todo);
      showInfo("todo create job started");
      reset();
      if ($closeOnCreate) {
        closeOffcanvas();
      }
    } catch (error) {
      showError(`todo create job failed\n${error}`);
    }
  }

  function closeOffcanvas() {
    showOffcanvas = false;
  }

  function toggleOffcanvas() {
    showOffcanvas = !showOffcanvas;
  }

  $: canCreate = $isValid && !$loading;
</script>

<button
  class="btn btn-primary rounded-pill px-3"
  type="button"
  on:click={toggleOffcanvas}
>
  <i class="bi bi-plus-lg" />
  New Todo
</button>

{#if showOffcanvas}
  <Offcanvas title="New Todo" on:close={closeOffcanvas}>
    <form
      spellcheck="false"
      autocomplete="off"
      on:submit|preventDefault={handleSubmit}
    >
      <div class="mb-3">
        <label for="name" class="form-label"
          >Name <span class="text-danger">*</span></label
        >
        <input
          type="text"
          class="form-control"
          id="name"
          placeholder="Name..."
          bind:value={$name}
        />
      </div>
      <div class="mb-3">
        <label for="description" class="form-label">Description</label>
        <textarea
          class="form-control"
          id="description"
          rows="5"
          bind:value={$description}
        />
      </div>
      <div class="mb-3">
        <div class="form-check">
          <input
            id="close-on-save-check"
            class="form-check-input"
            type="checkbox"
            bind:checked={$closeOnCreate}
          />
          <label class="form-check-label" for="close-on-save-check">
            Close on Create
          </label>
        </div>
      </div>
      <button
        type="submit"
        class="btn btn-primary rounded-pill px-3"
        disabled={!canCreate}
      >
        <i class="bi bi-plus-lg" />
        Create
      </button>
    </form>
  </Offcanvas>
{/if}
