import CreateUser from "@/components/CreateUser";
import UserList from "@/components/UserList";

// Pagina principal
// TODO - Implementar State Management entre componentes
// TODO - Implementar validação de campos
export default function Home() {
  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <h1 className="text-4xl font-bold text-center">Usuários</h1>

      <UserList />
      <h2 className="text-2xl font-bold">Criar Novo Usuário</h2>
      <CreateUser />
    </div>
  );
}
