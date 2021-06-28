import { writable } from "svelte/store";

let toastId = 1;

const TIMEOUT = 2500;

export const toasts = writable(/** @type {import("../types").Toast[]} */ ([]));

/**
 * @param {string} message
 */
export function showInfo(message) {
  showToast(message, "primary");
}

/**
 * @param {string} message
 */
export function showError(message) {
  showToast(message, "danger");
}

/**
 * @param {string} message
 * @param {"primary" | "danger"} type
 */
export function showToast(message, type = "primary") {
  toasts.update((t) => [
    ...t,
    {
      id: toastId++,
      message,
      type,
    },
  ]);

  setTimeout(() => toasts.update((t) => t.filter((_, i) => i)), TIMEOUT);
}
