//@ts-ignore
const keycloak = new Keycloak();

/**
 * @param {Function} cb
 */
export async function initAuth(cb) {
  try {
    const isAuthenticated = await keycloak.init({ pkceMethod: "S256" });
    if (!isAuthenticated) {
      keycloak.login();
    } else {
      cb();
    }
  } catch (error) {
    console.log(error);
  }
}

/**
 * @returns {string} Access Token
 */
export function accessToken() {
  return keycloak.token;
}

export function logout() {
  keycloak.logout();
}
