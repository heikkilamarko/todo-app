import { writable } from "svelte/store";

let nextId = 1;

const TIMEOUT = 2500;

/** @type {import("svelte/store").Writable<import("../types").Toast[]>} */
export const toasts = writable([]);

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
      id: nextId++,
      message,
      type,
    },
  ]);

  setTimeout(() => toasts.update((t) => t.filter((_, i) => i)), TIMEOUT);
}
