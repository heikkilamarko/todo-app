import { writable } from "svelte/store";

/**
 * @returns {import("../index").ToasterStore}
 */
export function createToasterStore() {
  let toastId = 1;

  const TIMEOUT = 2500;

  /** @type {import("../index").ToastsStore} */
  const toasts = writable([]);

  /**
   * @param {string} message
   */
  function showInfo(message) {
    showToast(message, "info");
  }

  /**
   * @param {string} message
   */
  function showError(message) {
    showToast(message, "error");
  }

  /**
   * @param {string} message
   * @param {import("../index").ToastType} type
   */
  function showToast(message, type = "info") {
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

  return {
    toasts,
    showInfo,
    showError,
    showToast,
  };
}
