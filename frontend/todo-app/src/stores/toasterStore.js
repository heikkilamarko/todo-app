import { writable } from "svelte/store";

let nextId = 0;

const TIMEOUT = 2500;

export const toasts = writable([]);

/**
 * @param {string} message
 */
export function showInfo(message) {
  showToast(message, "success");
}

/**
 * @param {string} message
 */
export function showError(message) {
  showToast(message, "danger");
}

/**
 * @param {string} message
 * @param {"success"|"danger"} type
 */
export function showToast(message, type = "success") {
  toasts.update((t) => [
    ...t,
    {
      id: ++nextId,
      message,
      type,
    },
  ]);

  setTimeout(() => toasts.update((t) => t.filter((_, i) => i)), TIMEOUT);
}
