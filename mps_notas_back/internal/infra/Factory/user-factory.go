package factory

import (
	"database/sql"
	"mps_notas_back/internal/infra/model"
	"mps_notas_back/internal/infra/repository"
)

type UserFactory interface {
	GenFactory() repository.UserRepository
}

type UserDAOInMemory struct{}

func (UserDAOInMemory) GenFactory() repository.UserRepositoryInMemory {
	return repository.UserRepositoryInMemory{Users: []model.UserDAO{}, CurrentID: 1}
}

type UserDAOSql struct{}

func (UserDAOSql) GenFactory(db *sql.DB) repository.UserRepository {
	return repository.UserRepositorySql{DB: db}
}
