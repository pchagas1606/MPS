package factory

import "mps_notas_back/internal/infra/model"

type UserFactory interface {
	GenFactory(string) *model.UserDAO
}
