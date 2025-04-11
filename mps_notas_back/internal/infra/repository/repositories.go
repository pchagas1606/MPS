package repository

import "mps_notas_back/internal/infra/model"

type UserRepository interface {
	FindAll() ([]model.UserDAO, error)
	FindByID(id int) (*model.UserDAO, error)
	FindByEmail(email string) (*model.UserDAO, error)
	Create(input model.NewUserInput) (model.UserDAO, error)
	Update(id int, input model.NewUserInput) (model.UserDAO, error)
	Delete(id int) error
}

type TaskRepository interface {
	FindAll() ([]model.TaskDAO, error)
	FindByID(id int) (*model.TaskDAO, error)
	Create(input model.NewTaskInput) (error)
	Update(id int, input model.NewTaskInput) (model.TaskDAO, error)
	Delete(id int) error
}
