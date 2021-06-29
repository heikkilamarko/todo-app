import { createToasterStore } from "todo-app-common";
import { createTodoStore } from "./todoStore";
import { createTodoFormStore } from "./todoFormStore";
import { createNotificationStore } from "./notificationStore";

/** @type {import("../types").Stores} */
export const stores = {
  toasterStore: null,
  todoStore: null,
  todoFormStore: null,
  notificationStore: null,
};

export function createStores() {
  stores.toasterStore = createToasterStore();
  stores.todoStore = createTodoStore();
  stores.todoFormStore = createTodoFormStore();
  stores.notificationStore = createNotificationStore(stores);
}
