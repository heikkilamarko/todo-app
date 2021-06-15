<script>
  import { name, description, canCreate, create } from "../stores/todoStore";
</script>

<button
  class="btn btn-primary rounded-pill px-3"
  type="button"
  data-bs-toggle="offcanvas"
  data-bs-target="#offcanvas-right"
  aria-controls="offcanvas-right"
>
  <i class="bi bi-plus-lg" />
  New Todo
</button>

<div
  id="offcanvas-right"
  class="offcanvas offcanvas-end"
  tabindex="-1"
  data-bs-scroll="true"
  data-bs-backdrop="false"
  aria-labelledby="offcanvas-right-label"
>
  <div class="offcanvas-header">
    <h5 id="offcanvas-right-label">New Todo</h5>
    <button
      type="button"
      class="btn-close text-reset"
      data-bs-dismiss="offcanvas"
      aria-label="Close"
    />
  </div>
  <div class="offcanvas-body">
    <form
      spellcheck="false"
      autocomplete="off"
      on:submit|preventDefault={create}
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
          rows="3"
          bind:value={$description}
        />
      </div>
      <button
        type="submit"
        class="btn btn-primary rounded-pill px-3"
        disabled={!$canCreate}
      >
        <i class="bi bi-plus-lg" />
        Create
      </button>
    </form>
  </div>
</div>
