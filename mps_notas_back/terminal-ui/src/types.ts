export interface Task {
  id: number;
  title: string;
  description: string;
  startDate: string;
  endDate: string;
  created_at: string;
}

export interface NewTaskInput {
  title: string;
  description: string;
  startDate: string;
  endDate: string;
}

export interface User {
  id: number;
  name: string;
  email: string;
  created_at: string;
}

export interface Credentials {
  email: string;
  password: string;
}
