import { writable } from "svelte/store";
import {
  HubConnection,
  HubConnectionBuilder,
  LogLevel,
} from "@microsoft/signalr";
import { config } from "../shared/config";
import { accessToken } from "../shared/auth";
import { Notification } from "../shared/utils";
import { showError } from "./toasterStore";
import { load } from "./todoStore";

export const connected = writable(false);

export async function connect() {
  const url = `${config.apiUrl}/push/notifications`;

  const connection = buildConnection(url);

  connection.onclose(() => connected.set(false));
  connection.onreconnecting(() => connected.set(false));
  connection.onreconnected(() => connected.set(true));

  connection.on(config.notificationMethod, async (notification) => {
    const { type, data } = notification ?? {};
    switch (type) {
      case Notification.TodoCreatedOk:
      case Notification.TodoCompletedOk:
        await load();
        break;
      case Notification.TodoCreatedError:
      case Notification.TodoCompletedError:
        showError(`error: ${data.code}\n${data.message || "-"}`);
        break;
    }
  });

  try {
    await connection.start();
    connected.set(true);
    return () => connection.stop();
  } catch (e) {
    showError(`real-time connection error\n${e}`);
  }
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
