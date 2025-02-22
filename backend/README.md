# MPS Backend

API REST desenvolvida com Hono.js e TypeScript para gerenciamento de usuários.

## Estrutura do Projeto

```
src/
  ├── controllers/     # Controladores da aplicação
  ├── models/         # Modelos e repositórios
  ├── routes/         # Definições de rotas
  ├── services/       # Lógica de negócios
  └── index.ts        # Ponto de entrada da aplicação
```

## Pré-requisitos

- Node.js (Bun)
- Bun (gerenciador de pacotes)

## Instalação

```sh
# Instalar dependências
bun install
```

## Executando o Projeto

```sh
# Iniciar o servidor de desenvolvimento
bun run dev
```

O servidor estará disponível em `http://localhost:3000`

## Endpoints da API

### Usuários

- `GET /api/users` - Lista todos os usuários
- `GET /api/users/:id` - Busca um usuário específico
- `POST /api/users` - Cria um novo usuário

### Monitoramento

- `GET /health` - Verifica o status da API

## Configuração CORS

O backend está configurado para aceitar requisições do frontend rodando em `http://localhost:3001`

## Tecnologias Utilizadas

- [Hono.js](https://hono.dev/) - Framework web
- TypeScript - Linguagem de programação
- Bun - Runtime JavaScript e gerenciador de pacotes

## Estrutura de Camadas

- **Controllers**: Manipulam as requisições HTTP
- **Services**: Contêm a lógica de negócios
- **Repositories**: Gerenciam o acesso aos dados
- **Routes**: Definem os endpoints da API

## Desenvolvimento

O projeto utiliza uma arquitetura em camadas com injeção de dependências para melhor manutenibilidade e testabilidade.

Para adicionar novos recursos, siga o padrão existente de:

1. Criar o modelo
2. Implementar o repositório
3. Criar o serviço
4. Adicionar o controller
5. Definir as rotas
