/// <reference types="svelte" />

import { Writable } from "svelte/store";

export * from "./components";
export * from "./stores";

export type ToastType = "info" | "error";

export interface Toast {
  id: number;
  type: ToastType;
  message: string;
}

export type ToastsStore = Writable<Toast[]>;

export interface ToasterStore {
  toasts: ToastsStore;
  showInfo: (message: string) => void;
  showError: (message: string) => void;
  showToast: (message: string, type: ToastType) => void;
}
