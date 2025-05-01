import { api } from "./auth";

export async function getAllUsers() {
  const { data } = await api.get("/users");
  return data;
}

export async function createUser(
  name: string,
  email: string,
  password: string
) {
  const { data } = await api.post("/users", { name, email, password });
  return data;
}
