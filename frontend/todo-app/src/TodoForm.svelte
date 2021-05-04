<script>
  import { createTodo } from "./utils";

  let name = "";
  let description = "";

  async function submit() {
    try {
      await createTodo({
        name,
        description: description || null,
      });
      name = "";
      description = "";
    } catch (e) {
      alert(`Todo creation failed: ${e}`);
    }
  }

  $: canSubmit = !!name;
</script>

<form spellcheck="false" autocomplete="off" on:submit|preventDefault={submit}>
  <div class="mb-3">
    <label for="name" class="form-label"
      >Name <span class="text-danger">*</span></label
    >
    <input
      type="text"
      class="form-control"
      id="name"
      placeholder="Name..."
      bind:value={name}
    />
  </div>
  <div class="mb-3">
    <label for="description" class="form-label">Description</label>
    <textarea
      class="form-control"
      id="description"
      rows="3"
      bind:value={description}
    />
  </div>
  <button type="submit" class="btn btn-success" disabled={!canSubmit}
    >Create</button
  >
</form>
