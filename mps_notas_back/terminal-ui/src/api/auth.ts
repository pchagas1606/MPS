import axios from "axios";
import dotenv from "dotenv";
dotenv.config();

export const api = axios.create({
  baseURL: process.env.API_URL,
});

let token: string | null = null;

export function setToken(t: string) {
  token = t;
  api.defaults.headers.common["Authorization"] = `Bearer ${token}`;
}

export async function login(email: string, password: string) {
  const response = await api.post("/login", { email, password });
  setToken(response.data.token);
  return response.data;
}
