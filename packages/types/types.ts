// Shared TypeScript type definitions
// These should match the Go types in the backend

export interface User {
  id: string;
  email: string;
  username: string;
  firstName: string;
  lastName: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateUserRequest {
  email: string;
  username: string;
  firstName: string;
  lastName: string;
  password: string;
}

export interface UpdateUserRequest {
  email?: string;
  firstName?: string;
  lastName?: string;
}

export interface ApiResponse<T> {
  data?: T;
  error?: string;
  message: string;
}

export interface HealthResponse {
  message: string;
  status: string;
}
