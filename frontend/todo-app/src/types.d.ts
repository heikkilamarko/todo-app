export interface Config {
  apiUrl: string;
  notificationMethod: string;
  auth: Keycloak.KeycloakConfig;
}

export interface Toast {
  id: number;
  type: "success" | "danger";
  message: string;
}

export type NotificationType =
  | "todo.created.ok"
  | "todo.created.error"
  | "todo.completed.ok"
  | "todo.completed.error";

export interface NewTodo {
  name: string;
  description?: string;
}

export interface ServerTodo {
  id: number;
  name: string;
  description?: string;
  created_at: string;
  updated_at: string;
}

export interface Todo {
  id: number;
  name: string;
  description?: string;
  created_at: Date;
  updated_at: Date;
}
