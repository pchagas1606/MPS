package service

import (
	"mps_notas_back/internal/infra/model"
	"mps_notas_back/internal/infra/repository"
)

// Task implementa a lógica de negócio relacionada as tarefas
type TaskService struct {
	repo repository.TaskRepository
}

// NewUserService cria uma nova instância do serviço de usuários
func NewTaskService(repo repository.TaskRepository ) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// GetAllUsers retorna todos os usuários
func (s *TaskService) GetAllTask() ([]model.TaskDAO, error) {
	return s.repo.FindAll()
}

// GetUserByID retorna um usuário pelo ID
func (s *TaskService) GetTaskByID(id int) (*model.TaskDAO, error) {
	return s.repo.FindByID(id)
}

// CreateUser cria um novo usuário
func (s *TaskService) CreateTask(input model.NewTaskInput) (error) {
	// Aqui poderia haver validações adicionais, prox atividade
	// if input.Name == "" || input.Email == "" {
	//     throw error
	// }
	return s.repo.Create(input)
}

// Update atualiza um usuário
func (s *TaskService) UpdateTask(id int, input model.NewTaskInput) (model.TaskDAO, error) {
	// Aqui poderia haver validações adicionais, prox atividade
	// if input.Name == "" || input.Email == "" {
	//     throw error
	// }
	return s.repo.Update(id, input)
}

// Update atualiza um usuário
func (s *TaskService) DeleteTask(id int) error {
	// Aqui poderia haver validações adicionais, prox atividade
	// if input.Name == "" || input.Email == "" {
	//     throw error
	// }
	return s.repo.Delete(id)
}
