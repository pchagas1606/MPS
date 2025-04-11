package repository

import (
	"errors"
	"mps_notas_back/internal/buisness/security"
	"mps_notas_back/internal/infra/model"
	"time"
)

// UserRepository implementa o acesso aos dados dos usuários
type UserRepositoryInMemory struct {
	Users     []model.UserDAO
	CurrentID int
}

// NewUserRepository cria uma nova instância do repositório de usuários

// FindAll retorna todos os usuários
func (r *UserRepositoryInMemory) FindAll() ([]model.UserDAO, error) {

	// Cria uma cópia da slice para evitar problemas de concorrência
	result := make([]model.UserDAO, len(r.Users))
	copy(result, r.Users)
	return result, nil
}

// FindByID retorna um usuário pelo ID ou nil se não for encontrado
func (r *UserRepositoryInMemory) FindByID(id int) (*model.UserDAO, error) {

	for i, user := range r.Users {
		if user.ID == id {
			// Retorna uma cópia para evitar problemas de concorrência
			userCopy := r.Users[i]
			return &userCopy, nil
		}
	}
	return nil, errors.New("error nao existe usuario com esse ID")
}

// FindByEmail retorna um usuário pelo Email ou nil se não for encontrado
func (r *UserRepositoryInMemory) FindByEmail(email string) (*model.UserDAO, error) {

	for i, user := range r.Users {
		if user.Email == email {
			// Retorna uma cópia para evitar problemas de concorrência
			userCopy := r.Users[i]
			return &userCopy, nil
		}
	}
	return nil, errors.New("error, nao existe usuario com esse email")
}

// Create adiciona um novo usuário e retorna o usuário criado, usando mutex para garantir a concorrência
func (r *UserRepositoryInMemory) Create(input model.NewUserInput) (model.UserDAO, error) {

	hashed_password, err := security.Hash(input.Password)
	if err != nil {
		return model.UserDAO{}, err
	}
	user := model.UserDAO{
		ID:            r.CurrentID,
		Name:          input.Name,
		Email:         input.Email,
		Password_Hash: string(hashed_password),
		CreatedAt:     time.Now(),
	}

	r.CurrentID++
	r.Users = append(r.Users, user)
	return user, nil
}

// Update atualiza um usuário existente e retorna o usuário atualizado
func (r *UserRepositoryInMemory) Update(id int, input model.NewUserInput) (model.UserDAO, error) {

	// Verificar se o ID foi fornecido
	if id == 0 {
		return model.UserDAO{}, errors.New("ID do usuário não fornecido")
	}

	// Encontrar o usuário pelo ID
	var found bool
	var updatedUser model.UserDAO

	for i, user := range r.Users {
		if user.ID == id {
			// Atualizar os campos fornecidos
			if input.Name != "" {
				r.Users[i].Name = input.Name
			}
			if input.Email != "" {
				r.Users[i].Email = input.Email
			}
			if input.Password != "" {
				hashedPassword, err := security.Hash(input.Password)
				if err != nil {
					return model.UserDAO{}, err
				}
				r.Users[i].Password_Hash = string(hashedPassword)
			}

			updatedUser = r.Users[i]
			found = true
			break
		}
	}

	if !found {
		return model.UserDAO{}, errors.New("usuário não encontrado")
	}

	return updatedUser, nil
}

// Delete remove um usuário pelo ID
func (r *UserRepositoryInMemory) Delete(id int) error {

	if id == 0 {
		return errors.New("ID do usuário não fornecido")
	}

	found := false
	for i, user := range r.Users {
		if user.ID == id {
			// Remover o usuário da slice
			r.Users = append(r.Users[:i], r.Users[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("usuário não encontrado")
	}

	return nil
}
