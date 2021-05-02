import axios from "axios";
import { HubConnection, HubConnectionBuilder } from "@microsoft/signalr";

const API_URL = import.meta.env.VITE_PUBLIC_API_URL;

export const NOTIFICATION_METHOD_NAME = import.meta.env
  .VITE_PUBLIC_NOTIFICATION_METHOD_NAME;

const api = axios.create();
// @ts-ignore
api.defaults.baseURL = API_URL;

/**
 * Get todos from server.
 * @returns {Promise<Array<import("./types").Notification>>}
 */
export async function getTodos() {
  try {
    var response = await api.get("/todos");
    /** @type {Array<import("./types").ServerNotification>} */
    var notifications = response?.data?.data ?? [];
    return notifications.map(toNotification);
  } catch (e) {
    console.error(e);
  }
}

/**
 * Get SignalR Hub connection.
 * @returns {HubConnection}
 */
export function getSignalRConnection() {
  return new HubConnectionBuilder()
    .withUrl(`${API_URL}/push/notifications`)
    .build();
}

/**
 * Maps a server notification to a client one.
 * @param {import("./types").ServerNotification} notification
 * @returns {import("./types").Notification}
 */
export function toNotification(notification) {
  return {
    ...notification,
    created_at: new Date(notification.created_at),
    updated_at: new Date(notification.updated_at),
  };
}
