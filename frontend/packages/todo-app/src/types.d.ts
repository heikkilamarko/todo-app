export interface Config {
  apiUrl: string;
  notificationMethod: string;
  auth: Keycloak.KeycloakConfig;
  profileUrl: string;
  dashboardUrl: string;
}

export interface Toast {
  id: number;
  type: "primary" | "danger";
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

export interface GetTodosRequest {
  offset: number;
  limit: number;
}

export interface GetTodosResponse {
  meta: GetTodosResponseMeta;
  data: ServerTodo[];
}

export interface GetTodosResponseMeta {
  offset: number;
  limit: number;
}
