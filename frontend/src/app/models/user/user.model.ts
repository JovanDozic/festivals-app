export interface User {
  id: number;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string | null;
  username: string;
  email: string;
  password: string;
  role: string;
}
