<script>
  import { fade } from "svelte/transition";
</script>

<section in:fade={{ delay: 500, duration: 300 }}>
  <h1 class="display-5 fw-light text-primary mt-5">
    It seems you have nothing to do!
  </h1>
  <h2 class="display-6 text-muted">
    ☝️ Add some todos by tapping the button above.
  </h2>
</section>
