<script>
  import { logout, userName } from "../shared/auth";
</script>

<button
  type="button"
  class="btn btn-outline-danger rounded-pill position-fixed px-3 m-2 top-0 end-0"
  on:click={logout}><i class="bi bi-power" /> Logout {userName()}</button
>
