import { writable } from "svelte/store";
import Centrifuge from "centrifuge";
import { config } from "../shared/config";
import * as api from "../shared/api";

/**
 * @param {import("../types").Stores} stores
 * @returns {import("../types").NotificationStore}
 */
export function createNotificationStore(stores) {
  const {
    toasterStore: { showError },
    todoStore: { getTodos },
  } = stores;

  const connected = writable(null);

  /**
   * @returns {Promise<() => void>}
   */
  async function connect() {
    let token;

    try {
      token = await api.getToken();
    } catch (error) {
      showError(`real-time connection error\n${error}`);
      return;
    }

    const centrifuge = new Centrifuge(config.notificationsUrl);
    centrifuge.setToken(token);

    centrifuge.on("connect", () => connected.set(true));
    centrifuge.on("disconnect", () => connected.set(false));

    centrifuge.subscribe("notifications", async (ctx) => {
      /** @type {{type: import("../types").NotificationType, data: any}} */
      const { type, data } = ctx.data ?? {};
      switch (type) {
        case "todo.create.ok":
        case "todo.complete.ok":
          try {
            await getTodos();
          } catch (error) {
            showError(`todo loading failed\n${error}`);
          }
          break;
        case "todo.create.error":
        case "todo.complete.error":
          showError(`error: ${data.code}\n${data.message || "-"}`);
          break;
      }
    });

    try {
      centrifuge.connect();
      return () => centrifuge.disconnect();
    } catch (error) {
      connected.set(false);
      showError(`real-time connection error\n${error}`);
    }
  }

  return {
    connected,
    connect,
  };
}
