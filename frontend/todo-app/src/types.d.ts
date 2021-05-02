export interface ServerNotification {
  id: number;
  name: string;
  description?: string;
  created_at: string;
  updated_at: string;
}

export interface Notification {
  id: number;
  name: string;
  description?: string;
  created_at: Date;
  updated_at: Date;
}
