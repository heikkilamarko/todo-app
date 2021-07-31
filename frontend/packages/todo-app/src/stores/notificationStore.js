import { writable } from "svelte/store";
import {
  HubConnection,
  HubConnectionBuilder,
  LogLevel,
} from "@microsoft/signalr";
import { config } from "../shared/config";
import { accessToken } from "../shared/auth";

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
   * @returns {Promise<() => void>} cleanup function
   */
  async function connect() {
    const url = `${config.apiUrl}/notifications`;

    const connection = buildConnection(url);

    connection.onclose(() => connected.set(false));
    connection.onreconnecting(() => connected.set(false));
    connection.onreconnected(() => connected.set(true));

    connection.on(config.notificationMethod, async (notification) => {
      /** @type {{type: import("../types").NotificationType, data: any}} */
      const { type, data } = notification ?? {};
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
      await connection.start();
      connected.set(true);
      return () => connection.stop();
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

/**
 * @param {string} url
 * @returns {HubConnection}
 */
function buildConnection(url) {
  return new HubConnectionBuilder()
    .withUrl(url, {
      accessTokenFactory: () => accessToken(),
    })
    .configureLogging(LogLevel.Critical)
    .withAutomaticReconnect({
      nextRetryDelayInMilliseconds: () => 5000,
    })
    .build();
}
