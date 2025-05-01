import { api } from "./auth";

export async function checkHealth() {
  const { data } = await api.get("/health");
  return data;
}
