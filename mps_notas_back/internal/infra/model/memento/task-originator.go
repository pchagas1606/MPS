package memento

import "mps_notas_back/internal/infra/model"

type TaskOriginator struct {
	currentState model.TaskDAO
}

func (o *TaskOriginator) SetState(state model.TaskDAO) {
	o.currentState = state
}

func (o *TaskOriginator) GetState() model.TaskDAO {
	return o.currentState
}

func (o *TaskOriginator) SaveToMemento() TaskMemento {
	return NewTaskMemento(o.currentState)
}

func (o *TaskOriginator) RestoreFromMemento(m TaskMemento) {
	o.currentState = m.GetSavedState()
}
