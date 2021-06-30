<script>
  import { Github, Grid, Person, PersonFill, Power } from "todo-app-common";
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
    <PersonFill />
    {userName()}
  </button>
  <ul class="dropdown-menu">
    <li>
      <a
        class="dropdown-item"
        href={config.profileUrl}
        target="_blank"
        rel="noreferrer"
      >
        <Person clazz="text-primary" />
        Profile
      </a>
    </li>
    <li>
      <a
        class="dropdown-item"
        href={config.dashboardUrl}
        target="_blank"
        rel="noreferrer"
      >
        <Grid clazz="text-primary" />
        Dashboard
      </a>
    </li>
    <li>
      <a
        class="dropdown-item"
        href="https://github.com/heikkilamarko/todo-app"
        target="_blank"
        rel="noreferrer"
      >
        <Github clazz="text-primary" />
        GitHub
      </a>
    </li>
    <li><hr class="dropdown-divider" /></li>
    <li>
      <a class="dropdown-item" href="/" on:click|preventDefault={logout}>
        <Power clazz="text-primary" />
        Logout
      </a>
    </li>
  </ul>
</div>

<style>
  .dropdown-toggle::after {
    content: none;
  }
</style>
