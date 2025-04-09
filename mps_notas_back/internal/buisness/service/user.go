package service

import (
	"errors"
	"mps_notas_back/internal/buisness/auth"

	"mps_notas_back/internal/buisness/security"
	"mps_notas_back/internal/infra/model"
	"mps_notas_back/internal/infra/repository"
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
func (s *UserService) GetAllUsers() []model.UserDAO {
	return s.repo.FindAll()
}

// GetUserByID retorna um usuário pelo ID
func (s *UserService) GetUserByID(id int) *model.UserDAO {
	return s.repo.FindByID(id)
}

// CreateUser cria um novo usuário
func (s *UserService) CreateUser(input model.NewUserInput) (model.UserDAO, error) {
	// Aqui poderia haver validações adicionais, prox atividade
	// if input.Name == "" || input.Email == "" {
	//     throw error
	// }
	return s.repo.Create(input)
}

// Login Recebe email e senha do usuario compara com o hash salvo se forem equivalentes retorna um token para o uso da API
func (s *UserService) Login(input model.AuthUserInput) (model.Token, error) {

	user := s.repo.FindByEmail(input.Email)
	if user == nil {
		return model.Token{}, errors.New("user does not exists")
	}
	// Verifica se a senha recebica equivale ao hash
	if err := security.VerifyPassword(input.Password, user.Password_Hash); err != nil {
		return model.Token{}, err
	}
	// Gera o token
	t, err := auth.CreateToken(user.ID)
	if err != nil {
		return model.Token{}, err
	}
	return model.Token{Token: t}, nil
}
