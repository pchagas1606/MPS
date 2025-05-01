import { api } from "./auth";

export async function getAllTasks() {
  const { data } = await api.get("/tasks");
  return data;
}

export async function getTaskById(id: number) {
  const { data } = await api.get(`/tasks/${id}`);
  return data;
}

export async function createTask(
  title: string,
  description: string,
  start: string,
  end: string
) {
  const { data } = await api.post("/tasks", {
    title,
    description,
    "start-date": start,
    "end-date": end,
  });
  return data;
}

export async function updateTask(
  id: number,
  title: string,
  description: string,
  start: string,
  end: string
) {
  const { data } = await api.put(`/tasks/${id}`, {
    title,
    description,
    "start-date": start,
    "end-date": end,
  });
  return data;
}

export async function deleteTask(id: number) {
  const { data } = await api.delete(`/tasks/${id}`);
  return data;
}

export async function undoTaskUpdate(id: number) {
  const { data } = await api.post(`/tasks/${id}/undo`);
  return data;
}
