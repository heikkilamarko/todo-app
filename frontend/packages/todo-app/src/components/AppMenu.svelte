<script>
  import { config } from "../shared/config";
  import { logout, userName } from "../shared/auth";
</script>

<div class="btn-group position-fixed m-2 top-0 end-0">
  <button
    type="button"
    class="btn btn-outline-primary dropdown-toggle rounded-pill px-3"
    data-bs-toggle="dropdown"
    aria-expanded="false"
  >
    <i class="bi bi-person-fill" />
    {userName()}
  </button>
  <ul class="dropdown-menu">
    <li>
      <a class="dropdown-item" href={config.profileUrl} target="_blank"
        ><i class="bi bi-person text-primary" /> Profile</a
      >
    </li>
    <li>
      <a class="dropdown-item" href={config.dashboardUrl} target="_blank"
        ><i class="bi bi-grid text-primary" /> Dashboard</a
      >
    </li>
    <li>
      <a
        class="dropdown-item"
        href="https://github.com/heikkilamarko/todo-app"
        target="_blank"><i class="bi bi-github text-primary" /> GitHub</a
      >
    </li>
    <li><hr class="dropdown-divider" /></li>
    <li>
      <a class="dropdown-item" href="/" on:click|preventDefault={logout}
        ><i class="bi bi-power text-primary" /> Logout</a
      >
    </li>
  </ul>
</div>

<style>
  .dropdown-toggle::after {
    content: none;
  }
</style>
