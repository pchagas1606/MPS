package factory

import (
	"database/sql"
	"mps_notas_back/internal/infra/repository"
)

type TaskFactory interface {
	GenFactory() repository.UserRepository
}

// type UserDAOInMemory struct{}

// func (UserDAOInMemory) GenFactory() repository.UserRepositoryInMemory {
// 	return repository.UserRepositoryInMemory{Users: []model.UserDAO{}, CurrentID: 1}
// }

type TaskDAOSql struct{}

func (TaskDAOSql) GenFactory(db *sql.DB) repository.TaskRepository {
	return repository.TaskRepositorySql{DB: db}
}
