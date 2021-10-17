<script>
  import GithubIcon from "bootstrap-icons/icons/github.svg";
  import GridIcon from "bootstrap-icons/icons/grid.svg";
  import PersonIcon from "bootstrap-icons/icons/person.svg";
  import PersonFillIcon from "bootstrap-icons/icons/person-fill.svg";
  import PowerIcon from "bootstrap-icons/icons/power.svg";
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
    <PersonFillIcon />
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
        <PersonIcon />
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
        <GridIcon />
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
        <GithubIcon />
        GitHub
      </a>
    </li>
    <li><hr class="dropdown-divider" /></li>
    <li>
      <a class="dropdown-item" href="/" on:click|preventDefault={logout}>
        <PowerIcon />
        Logout
      </a>
    </li>
  </ul>
</div>

<style>
  .dropdown-toggle::after {
    content: none;
  }

  .dropdown-item :global(svg) {
    color: var(--bs-primary);
  }
</style>
