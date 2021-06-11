import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap-icons/font/bootstrap-icons.css";
import { load as loadConfig } from "./shared/config";
import { init as initAuth } from "./shared/auth";
import App from "./components/App.svelte";

async function bootstrap() {
  try {
    await loadConfig();
    const isAuthenticated = await initAuth();
    isAuthenticated &&
      new App({
        target: document.getElementById("app"),
      });
  } catch (error) {
    console.log(error);
  }
}

bootstrap();
