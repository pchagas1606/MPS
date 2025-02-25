package service

import (
	"mps_notas_back/internal/model"
	"mps_notas_back/internal/repository"
)

// UserService implementa a lógica de negócio relacionada aos usuários
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService cria uma nova instância do serviço de usuários
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// GetAllUsers retorna todos os usuários
func (s *UserService) GetAllUsers() []model.User {
	return s.repo.FindAll()
}

// GetUserByID retorna um usuário pelo ID
func (s *UserService) GetUserByID(id int) *model.User {
	return s.repo.FindByID(id)
}

// CreateUser cria um novo usuário
func (s *UserService) CreateUser(input model.NewUserInput) model.User {
	// Aqui poderia haver validações adicionais, prox atividade
	// if input.Name == "" || input.Email == "" {
	//     throw error
	// }
	
	return s.repo.Create(input)
}