import "bootstrap/js/dist/dropdown";
import "bootstrap/js/dist/offcanvas";
import "bootstrap-icons/font/bootstrap-icons.css";
import "./main.scss";
import { load as loadConfig } from "./shared/config";
import { init as initAuth } from "./shared/auth";
import App from "./components/App.svelte";

async function run() {
  try {
    await loadConfig();
    const isAuthenticated = await initAuth();
    isAuthenticated &&
      new App({
        target: document.getElementById("app"),
      });
  } catch (error) {
    console.log(error);
    renderError(error);
  }
}

/**
 * @param {Error} error
 */
function renderError(error) {
  document.body.innerHTML = `<main class="px-4 py-5 overflow-auto d-flex flex-column align-items-center vh-100 bg-danger text-white"><h1 class="display-1 fw-lighter">Close, but No Cigar</h1><p class="pt-2">${error}</p></main>`;
}

run();
