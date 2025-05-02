import { api } from "./auth";

export async function checkStatus() {
  const { data } = await api.get("/status");
  return data;
}
