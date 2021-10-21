import { config } from "./config";

/** @type {Keycloak.KeycloakInstance} */
let keycloak;

/**
 * @returns {Promise<boolean>}
 */
export async function init() {
  try {
    const { default: Keycloak } = await import("keycloak-js");

    keycloak = Keycloak(config.auth);
    const isAuthenticated = await keycloak.init({
      pkceMethod: "S256",
      // enableLogging: true,
    });
    if (isAuthenticated) {
      await keycloak.loadUserInfo();
    } else {
      await keycloak.login();
    }
    return isAuthenticated;
  } catch (error) {
    console.log(error);
    throw new Error("auth init failed");
  }
}

/**
 * @returns {Promise<string>} Access Token
 */
export async function accessToken() {
  try {
    await keycloak.updateToken(null);
    return keycloak.token;
  } catch (error) {
    await keycloak.login();
  }
}

/**
 * @returns {Promise}
 */
export async function logout() {
  await keycloak.logout();
}

/**
 * @returns {string}
 */
export function userName() {
  // @ts-ignore
  return keycloak.userInfo?.name ?? "<unknown user>";
}

export function isUserRole() {
  return isInRole("todo-user");
}

export function isViewerRole() {
  return isInRole("todo-viewer");
}

/**
 * @param {string} role
 * @returns {boolean}
 */
export function isInRole(role) {
  return keycloak.hasResourceRole(role, "todo-api");
}
