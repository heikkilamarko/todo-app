export interface Config {
  apiUrl: string;
  notificationMethod: string;
  auth: Keycloak.KeycloakConfig;
}

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
