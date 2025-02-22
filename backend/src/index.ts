import { Hono } from "hono";
import { logger } from "hono/logger";
import { prettyJSON } from "hono/pretty-json";
import { UserRepository } from "./models/UserRepository";
import { UserService } from "./services/UserService";
import { UserController } from "./controllers/UserController";
import { userRoutes } from "./routes/userRoutes";
import { cors } from "hono/cors";

// Inicializa a aplicação Hono
const app = new Hono();

// Configuração de Middleware
// Habilita CORS para permitir requisições de origens diferentes, frontend por exemplo
app.use(
  "*",
  cors({
    origin: ["http://localhost:3001"], // URL do frontend
    allowMethods: ["GET", "POST", "PUT", "DELETE", "OPTIONS"],
    allowHeaders: ["Content-Type", "Authorization"],
    exposeHeaders: ["Content-Length"],
    maxAge: 600,
    credentials: true,
  })
);

// Logger para registrar todas as requisições
app.use("*", logger());
// Formatador JSON para melhor legibilidade das respostas
app.use("*", prettyJSON());

// Configuração da Injeção de Dependências
// Cria instâncias das camadas da aplicação
const userRepository = new UserRepository();
const userService = new UserService(userRepository);
const userController = new UserController(userService);

// Definição das Rotas
// Rota de verificação de saúde da API
app.get("/health", (c) => {
  return c.json({ status: "ok" });
});

// Agrupa todas as rotas de usuário sob /api/users
app.route("/api/users", userRoutes(userController));

// Configuração do servidor
export default {
  port: 3000, // Porta em que o servidor irá rodar
  fetch: app.fetch, // Handler principal da aplicação
};
