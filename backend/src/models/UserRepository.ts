import type { User } from "./User";

// TODO - Banco de dados
// TODO - Implementar métodos de CRUD
export class UserRepository {
  private users: User[] = [];
  private currentId = 1;

  findAll(): User[] {
    return this.users;
  }

  findById(id: number): User | undefined {
    return this.users.find((user) => user.id === id);
  }

  create(userData: Omit<User, "id" | "createdAt">): User {
    const user: User = {
      id: this.currentId++,
      ...userData,
      createdAt: new Date(),
    };
    this.users.push(user);
    return user;
  }
}
