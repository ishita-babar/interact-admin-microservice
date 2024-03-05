export interface Log {
  id: string;
  level: string;
  title: string;
  description: string;
  path: string;
  timestamp: Date;
}

export interface User {
  id: string;
  username: string;
  role: string;
}
