//@ts-ignore
const keycloak = new Keycloak();

/**
 * @param {Function} cb
 */
export async function initAuth(cb) {
  try {
    const isAuthenticated = await keycloak.init({
      pkceMethod: "S256",
      // enableLogging: true,
    });
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
 * @returns {Promise<string>} Access Token
 */
export async function accessToken() {
  try {
    await keycloak.updateToken();
    return keycloak.token;
  } catch (error) {
    keycloak.login();
  }
}

export function logout() {
  keycloak.logout();
}
