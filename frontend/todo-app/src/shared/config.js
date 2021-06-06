import axios from "axios";
import { showError } from "./utils";

/** @type {import("../types").Config} */
export const config = {
  apiUrl: null,
  notificationMethod: null,
  auth: null,
};

export async function load() {
  try {
    import.meta.env.DEV ? loadDev() : await loadProd();
  } catch (error) {
    showError(`config loading failed\n${error}`);
  }
}

async function loadProd() {
  const { data } = await axios.get("/config.json");
  Object.assign(config, data);
}

function loadDev() {
  Object.assign(config, {
    apiUrl: import.meta.env.VITE_PUBLIC_API_URL,
    notificationMethod: import.meta.env.VITE_PUBLIC_NOTIFICATION_METHOD,
    auth: {
      url: import.meta.env.VITE_PUBLIC_AUTH_URL,
      realm: import.meta.env.VITE_PUBLIC_AUTH_REALM,
      clientId: import.meta.env.VITE_PUBLIC_AUTH_CLIENT_ID,
    },
  });
}
