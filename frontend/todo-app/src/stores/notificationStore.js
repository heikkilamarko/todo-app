import { writable, get } from "svelte/store";
import { getSignalRConnection, showError } from "../utils";
import { Notification } from "../constants";
import { config } from "./configStore";
import { load } from "./todoStore";

export const connected = writable(false);

export async function connect() {
  const c = get(config);

  let connection = getSignalRConnection(`${c.apiUrl}/push/notifications`);

  connection.onclose(() => connected.set(false));
  connection.onreconnecting(() => connected.set(false));
  connection.onreconnected(() => connected.set(true));

  connection.on(c.notificationMethodName, async (notification) => {
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
