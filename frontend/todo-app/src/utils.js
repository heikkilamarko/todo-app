import axios from "axios";
import { HubConnectionBuilder } from "@microsoft/signalr";

export const SIGNALR_RECEIVE_NOTIFICATION = "ReceiveNotification";

export async function getTodos() {
  try {
    var response = await axios.get("http://localhost:8080/todos");
    return response?.data?.data ?? [];
  } catch (e) {
    console.error(e);
  }
}

export function getSignalRConnection() {
  return new HubConnectionBuilder()
    .withUrl("http://localhost:8080/push/notifications")
    .build();
}
