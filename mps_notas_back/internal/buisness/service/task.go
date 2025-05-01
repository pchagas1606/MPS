package service

import (
	"errors"
	"mps_notas_back/internal/infra/model"
	"mps_notas_back/internal/infra/model/memento"
	"mps_notas_back/internal/infra/repository"
)

// Task implementa a lógica de negócio relacionada as tarefas
type TaskService struct {
	repo      repository.TaskRepository
	origin    memento.TaskOriginator
	caretaker memento.Caretaker
}

// NewUserService cria uma nova instância do serviço de usuários
func NewTaskService(repo repository.TaskRepository) *TaskService {
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
func (s *TaskService) CreateTask(input model.NewTaskInput) error {
	// Aqui poderia haver validações adicionais, prox atividade
	// if input.Name == "" || input.Email == "" {
	//     throw error
	// }
	return s.repo.Create(input)
}

// Update atualiza um usuário
func (s *TaskService) UpdateTask(id int, input model.NewTaskInput) (model.TaskDAO, error) {
	// get current state
	currentTask, err := s.repo.FindByID(id)
	if err != nil || currentTask == nil {
		return model.TaskDAO{}, errors.New("task not found")
	}

	// Save current state as memento
	s.origin.SetState(*currentTask)
	s.caretaker.AddMemento(s.origin.SaveToMemento())

	// Update task
	updatedTask, err := s.repo.Update(id, input)
	if err != nil {
		return model.TaskDAO{}, err
	}

	return updatedTask, nil
}

func (s *TaskService) UndoLastUpdate(id int) (model.TaskDAO, error) {
	m := s.caretaker.Undo()
	if m == nil {
		return model.TaskDAO{}, errors.New("no operation to undo")
	}

	previousState := m.GetSavedState()
	// Revert in DB
	input := model.NewTaskInput{
		Title:       previousState.Title,
		Description: previousState.Description,
		StartDate:   previousState.StartDate,
		EndDate:     previousState.EndDate,
	}

	return s.repo.Update(previousState.ID, input)
}

// Update atualiza um usuário
func (s *TaskService) DeleteTask(id int) error {
	// Aqui poderia haver validações adicionais, prox atividade
	// if input.Name == "" || input.Email == "" {
	//     throw error
	// }
	return s.repo.Delete(id)
}
