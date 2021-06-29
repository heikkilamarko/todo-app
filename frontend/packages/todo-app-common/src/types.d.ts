/// <reference types="svelte" />

import { Writable } from "svelte/store";

type ToastType = "info" | "error";

interface Toast {
  id: number;
  type: ToastType;
  message: string;
}

type ToastsStore = Writable<Toast[]>;

interface ToasterStore {
  toasts: ToastsStore;
  showInfo: (message: string) => void;
  showError: (message: string) => void;
  showToast: (message: string, type: ToastType) => void;
}
