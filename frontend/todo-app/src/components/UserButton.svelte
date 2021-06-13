<script>
  import { config } from "../shared/config";
  import { logout, userName } from "../shared/auth";
</script>

<div class="btn-group position-fixed m-2 top-0 end-0">
  <button
    type="button"
    class="btn btn-outline-success dropdown-toggle rounded-pill px-3"
    data-bs-toggle="dropdown"
    aria-expanded="false"
  >
    <i class="bi bi-person-fill" />
    {userName()}
  </button>
  <ul class="dropdown-menu">
    <li>
      <a class="dropdown-item" href={config.profileUrl} target="_blank"
        ><i class="bi bi-person text-success" /> Profile</a
      >
    </li>
    <li><hr class="dropdown-divider" /></li>
    <li>
      <a class="dropdown-item" href="/" on:click|preventDefault={logout}
        ><i class="bi bi-power text-success" /> Logout</a
      >
    </li>
  </ul>
</div>

<style>
  .dropdown-toggle::after {
    content: none;
  }
</style>
