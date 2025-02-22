import type { User } from "../models/User";
import { UserRepository } from "../models/UserRepository";

export class UserService {
  constructor(private userRepository: UserRepository) {}

  // Obtém todos os usuários do repositório
  getAllUsers(): User[] {
    return this.userRepository.findAll();
  }

  // Busca um usuário específico pelo ID
  getUserById(id: number): User | undefined {
    return this.userRepository.findById(id);
  }

  // Cria um novo usuário
  createUser(userData: Omit<User, "id" | "createdAt">): User {
    // Validação básica dos campos obrigatórios
    // if (!userData.name || !userData.email) {
    // throw new Error("Name and email are required"); // prox atividade
    // }
    return this.userRepository.create(userData);
  }
}
