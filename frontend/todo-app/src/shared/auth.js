import Keycloak from "keycloak-js";
import { config } from "./config";

/** @type {Keycloak.KeycloakInstance} */
let keycloak;

/**
 * @returns {Promise<boolean>}
 */
export async function init() {
  try {
    keycloak = Keycloak(config.auth);
    const isAuthenticated = await keycloak.init({
      pkceMethod: "S256",
      // enableLogging: true,
    });
    if (!isAuthenticated) {
      await keycloak.login();
    }
    return isAuthenticated;
  } catch (error) {
    console.log(error);
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
