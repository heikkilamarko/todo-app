import { writable } from "svelte/store";
import axios from "axios";
import { showError } from "../utils";

export const config = writable(null);
export const loading = writable(false);

export async function load() {
  import.meta.env.DEV ? loadDev() : await loadProd();
}

async function loadProd() {
  try {
    loading.set(true);
    var r = await axios.get("/config.json");
    config.set(r.data);
  } catch (e) {
    showError(`config loading failed\n${e}`);
  } finally {
    loading.set(false);
  }
}

function loadDev() {
  try {
    const c = {
      apiUrl: import.meta.env.VITE_PUBLIC_API_URL,
      notificationMethodName: import.meta.env
        .VITE_PUBLIC_NOTIFICATION_METHOD_NAME,
    };
    config.set(c);
  } catch (e) {
    showError(`config loading failed\n${e}`);
  }
}
