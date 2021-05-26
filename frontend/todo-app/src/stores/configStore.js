import { writable } from "svelte/store";
import axios from "axios";
import { showError } from "../utils";

export const config = writable(null);
export const loading = writable(false);

export async function load() {
  try {
    loading.set(true);
    var r = await axios.get("/config.json");
    config.set(r.data);
  } catch (e) {
    showError(`config loading failed\n${e}`);
  } finally {
    loading.set(false);
  }
}
