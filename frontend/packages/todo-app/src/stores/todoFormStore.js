import { derived, writable } from "svelte/store";

export const name = writable("");
export const description = writable("");

export const closeOnCreate = writable(true);

export const todo = derived([name, description], ([$name, $description]) => ({
  name: $name,
  description: $description || null,
}));

export const isValid = derived(name, ($name) => !!$name);

export function reset() {
  name.set("");
  description.set("");
}
