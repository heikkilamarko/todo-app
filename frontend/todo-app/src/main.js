import { load as loadConfig } from "./shared/config";
import { init as initAuth } from "./shared/auth";
import App from "./components/App.svelte";

async function bootstrap() {
  await loadConfig();
  const isAuthenticated = await initAuth();
  isAuthenticated &&
    new App({
      target: document.getElementById("app"),
    });
}

bootstrap();
