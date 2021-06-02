<script>
  import { logout } from "../shared/auth";
</script>

<button
  type="button"
  class="btn btn-danger position-fixed m-2 top-0 end-0 fw-bold"
  on:click={logout}>Logout</button
>
