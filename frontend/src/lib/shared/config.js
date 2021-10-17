import ky from "ky";

/** @type {import("../../types").Config} */
export const config = {
  apiUrl: null,
  notificationsUrl: null,
  auth: null,
  profileUrl: null,
  dashboardUrl: null,
};

export async function load() {
  try {
    import.meta.env.DEV ? loadDev() : await loadProd();
  } catch (error) {
    console.log(error);
    throw new Error("config loading failed");
  }
}

async function loadProd() {
  Object.assign(config, await ky.get("/config.json").json());
}

function loadDev() {
  Object.assign(config, {
    apiUrl: import.meta.env.VITE_PUBLIC_API_URL,
    notificationsUrl: import.meta.env.VITE_PUBLIC_NOTIFICATIONS_URL,
    auth: {
      url: import.meta.env.VITE_PUBLIC_AUTH_URL,
      realm: import.meta.env.VITE_PUBLIC_AUTH_REALM,
      clientId: import.meta.env.VITE_PUBLIC_AUTH_CLIENT_ID,
    },
    profileUrl: import.meta.env.VITE_PUBLIC_PROFILE_URL,
    dashboardUrl: import.meta.env.VITE_PUBLIC_DASHBOARD_URL,
  });
}
