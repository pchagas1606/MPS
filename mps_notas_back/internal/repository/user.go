package repository

import (
	"mps_notas_back/internal/model"
	"mps_notas_back/internal/security"
	"sync"
	"time"
)

// UserRepository implementa o acesso aos dados dos usuários
type UserRepository struct {
	users     []model.User
	currentID int
	mutex     sync.RWMutex
}

// NewUserRepository cria uma nova instância do repositório de usuários
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:     []model.User{},
		currentID: 1,
	}
}

// FindAll retorna todos os usuários
func (r *UserRepository) FindAll() []model.User {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// Cria uma cópia da slice para evitar problemas de concorrência
	result := make([]model.User, len(r.users))
	copy(result, r.users)
	return result
}

// FindByID retorna um usuário pelo ID ou nil se não for encontrado
func (r *UserRepository) FindByID(id int) *model.User {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for i, user := range r.users {
		if user.ID == id {
			// Retorna uma cópia para evitar problemas de concorrência
			userCopy := r.users[i]
			return &userCopy
		}
	}
	return nil
}

// FindByEmail retorna um usuário pelo Email ou nil se não for encontrado
func (r *UserRepository) FindByEmail(email string) *model.User {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for i, user := range r.users {
		if user.Email == email {
			// Retorna uma cópia para evitar problemas de concorrência
			userCopy := r.users[i]
			return &userCopy
		}
	}
	return nil
}

// Create adiciona um novo usuário e retorna o usuário criado, usando mutex para garantir a concorrência
func (r *UserRepository) Create(input model.NewUserInput) (model.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	hashed_password, err := security.Hash(input.Password)
	if err != nil {
		return model.User{}, err
	}
	user := model.User{
		ID:            r.currentID,
		Name:          input.Name,
		Email:         input.Email,
		Password_Hash: string(hashed_password),
		CreatedAt:     time.Now(),
	}

	r.currentID++
	r.users = append(r.users, user)
	return user, nil
}
