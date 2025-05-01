package memento

type Caretaker struct {
	mementoStack []TaskMemento
}

func (c *Caretaker) AddMemento(m TaskMemento) {
	c.mementoStack = append(c.mementoStack, m)
}

func (c *Caretaker) Undo() *TaskMemento {
	if len(c.mementoStack) == 0 {
		return nil
	}
	// Pop last memento
	lastIndex := len(c.mementoStack) - 1
	m := c.mementoStack[lastIndex]
	c.mementoStack = c.mementoStack[:lastIndex]
	return &m
}
