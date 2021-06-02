import { initAuth } from "./shared/auth";
import App from "./components/App.svelte";

initAuth(
  () =>
    new App({
      target: document.getElementById("app"),
    })
);
