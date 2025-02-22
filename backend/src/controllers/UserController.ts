import type { Context } from "hono";
import { UserService } from "../services/UserService";

export class UserController {
  constructor(private userService: UserService) {}

  // Busca todos os usuários
  async index(c: Context) {
    const users = this.userService.getAllUsers();
    return c.json(users);
  }

  // Busca um usuário específico pelo ID
  async show(c: Context) {
    const id = parseInt(c.req.param("id"));
    const user = this.userService.getUserById(id);

    // Retorna 404 se o usuário não for encontrado
    if (!user) {
      return c.json({ error: "User not found" }, 404);
    }

    return c.json(user);
  }

  // Cria um novo usuário
  async create(c: Context) {
    try {
      const body = await c.req.json();

      // Validação básica dos campos obrigatórios, prox atividade
      // if (!body.name || !body.email) {
      //   return c.json({ error: "Name and email are required" }, 400);
      // }

      const user = this.userService.createUser({
        name: body.name,
        email: body.email,
      });

      return c.json(user, 201);
    } catch (error) {
      return c.json({ error: "Invalid request body" }, 400);
    }
  }
}
