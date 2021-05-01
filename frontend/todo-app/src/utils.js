import axios from "axios";
import { HubConnectionBuilder } from "@microsoft/signalr";

const API_URL = import.meta.env.VITE_PUBLIC_API_URL;

export const NOTIFICATION_METHOD_NAME = import.meta.env
  .VITE_PUBLIC_NOTIFICATION_METHOD_NAME;

const api = axios.create();
// @ts-ignore
api.defaults.baseURL = API_URL;

export async function getTodos() {
  try {
    var response = await api.get("/todos");
    return response?.data?.data ?? [];
  } catch (e) {
    console.error(e);
  }
}

export function getSignalRConnection() {
  return new HubConnectionBuilder()
    .withUrl(`${API_URL}/push/notifications`)
    .build();
}
