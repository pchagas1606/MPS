import type { User } from "../types/user";

const API_URL = "http://localhost:3000/api/users";

export const getAllUsers = async (): Promise<User[]> => {
  const response = await fetch(API_URL);
  if (!response.ok) {
    // throw new Error("Erro ao buscar usuários");
    console.error("Erro ao buscar usuários " + response.statusText);
  }
  return response.json();
};

export const getUserById = async (id: number): Promise<User | null> => {
  const response = await fetch(`${API_URL}/${id}`);
  if (!response.ok) {
    return null;
  }
  return response.json();
};

export const createUser = async (
  userData: Omit<User, "id" | "createdAt">
): Promise<User> => {
  const response = await fetch(API_URL, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(userData),
  });
  if (!response.ok) {
    // throw new Error("Erro ao criar usuário");
    console.error("Erro ao criar usuário " + response.statusText);
  }
  return response.json();
};
