package memento

import "mps_notas_back/internal/infra/model"

type TaskMemento struct {
	State model.TaskDAO
}

func NewTaskMemento(state model.TaskDAO) TaskMemento {
	return TaskMemento{State: state}
}

func (m TaskMemento) GetSavedState() model.TaskDAO {
	return m.State
}
