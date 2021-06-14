import { writable } from "svelte/store";

let nextId = 0;

const TIMEOUT = 2500;

/** @type {import("svelte/store").Writable<import("../types").Toast[]>} */
export const toasts = writable([]);

/**
 * Shows an info message
 * @param {string} message message
 */
export function showInfo(message) {
  showToast(message, "primary");
}

/**
 * Shows an error message
 * @param {string} message message
 */
export function showError(message) {
  showToast(message, "danger");
}

/**
 * Shows a message
 * @param {string} message message
 * @param {"primary" | "danger"} type message type
 */
export function showToast(message, type = "primary") {
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
