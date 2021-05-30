import { initAuth } from "./auth";
import App from "./App.svelte";

initAuth(
  () =>
    new App({
      target: document.getElementById("app"),
    })
);
