import { Hono } from "hono";
import { UserController } from "../controllers/UserController";

export function userRoutes(userController: UserController) {
  const router = new Hono();

  // Define as rotas para as operações de usuário
  router.get("/", (c) => userController.index(c)); // Lista todos os usuários
  router.get("/:id", (c) => userController.show(c)); // Busca um usuário específico
  router.post("/", (c) => userController.create(c)); // Cria um novo usuário

  return router;
}
