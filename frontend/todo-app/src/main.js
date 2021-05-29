import App from "./App.svelte";

async function run() {
  try {
    //@ts-ignore
    const keycloak = new Keycloak();
    const isAuthenticated = await keycloak.init({ pkceMethod: "S256" });
    if (!isAuthenticated) {
      keycloak.login();
    } else {
      //@ts-ignore
      window.accessToken = keycloak.token;
      new App({
        target: document.getElementById("app"),
      });
    }
  } catch (error) {
    console.log(error);
  }
}

run();
